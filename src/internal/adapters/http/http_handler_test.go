package http_test

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	httpadapter "github.com/vmlellis/port-sync/src/internal/adapters/http"
	"github.com/vmlellis/port-sync/src/internal/adapters/processor"
	"github.com/vmlellis/port-sync/src/internal/adapters/storage"
	"github.com/vmlellis/port-sync/src/internal/domain/entity"
	"github.com/vmlellis/port-sync/src/internal/domain/service"
)

func TestGetPortHandler(t *testing.T) {

	store := storage.NewMemoryStore()
	portService := service.NewPortService(store)
	processorService := processor.NewParallelProcessor(portService, processor.ParallelProcessorOpts{})
	handler := httpadapter.NewPortHandler(portService, processorService)

	port := &entity.Port{
		ID:       "PORT1",
		Name:     "Test Port",
		City:     "Test City",
		Country:  "Test Country",
		Province: "Test Province",
	}
	portService.SavePort(port)

	req, err := http.NewRequest("GET", "/ports/PORT1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetPort(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Test Port")
}

func TestGetPortHandler_NotFound(t *testing.T) {
	store := storage.NewMemoryStore()
	portService := service.NewPortService(store)
	processorService := processor.NewParallelProcessor(portService, processor.ParallelProcessorOpts{})
	handler := httpadapter.NewPortHandler(portService, processorService)

	req, err := http.NewRequest("GET", "/ports/INVALID_PORT", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetPort(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestBulkUploadPortsHandler(t *testing.T) {
	store := storage.NewMemoryStore()
	portService := service.NewPortService(store)
	processorService := processor.NewParallelProcessor(portService, processor.ParallelProcessorOpts{})
	handler := httpadapter.NewPortHandler(portService, processorService)

	jsonPayload := `{
			"PORT1": {
					"name": "Port One",
					"city": "City One",
					"country": "Country One",
					"coordinates": [10.10, 20.20],
					"province": "Province One",
					"timezone": "UTC+1",
					"unlocs": ["UNLOC1"],
					"code": "1234"
			},
			"PORT2": {
					"name": "Port Two",
					"city": "City Two",
					"country": "Country Two",
					"coordinates": [30.30, 40.40],
					"province": "Province Two",
					"timezone": "UTC+2",
					"unlocs": ["UNLOC2"],
					"code": "5678"
			}
	}`

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Create a form field for the file
	fileWriter, err := writer.CreateFormFile("file", "ports.json")
	assert.NoError(t, err)

	_, err = fileWriter.Write([]byte(jsonPayload))
	assert.NoError(t, err)

	writer.Close()

	req, err := http.NewRequest("POST", "/ports/bulk", &requestBody)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rr := httptest.NewRecorder()
	handler.BulkUploadPorts(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Contains(t, rr.Body.String(), "Ports uploaded successfully")
}
