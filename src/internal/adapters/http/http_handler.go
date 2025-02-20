package http

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/vmlellis/port-sync/src/internal/domain/contract"
)

// PortHandler provides HTTP endpoints to interact with port data.
type PortHandler struct {
	service   contract.PortService
	processor contract.ProcessorService
}

// NewPortHandler creates a new instance of PortHandler.
func NewPortHandler(service contract.PortService, processor contract.ProcessorService) *PortHandler {
	return &PortHandler{service: service, processor: processor}
}

// GetPort retrieves a port by its ID.
func (h *PortHandler) GetPort(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		JSONError(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 || pathParts[1] != "ports" {
		JSONError(w, "Invalid request path", http.StatusBadRequest)
		return
	}
	id := pathParts[2]

	port, found := h.service.GetPort(id)
	if !found {
		JSONError(w, "Port not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(port)
}

// BulkUploadPorts handles bulk upload of ports from a JSON file.
func (h *PortHandler) BulkUploadPorts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JSONError(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(50 * 1024 * 1024); err != nil { // Limit: 25MB
		JSONError(w, "File too large", http.StatusRequestEntityTooLarge)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		JSONError(w, "Invalid file upload", http.StatusBadRequest)
		return
	}
	defer file.Close()

	err = h.processor.Process(file)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CreatedResponse{Message: "Ports uploaded successfully"})
}

// RegisterRoutes sets up the HTTP routes for port operations.
func (h *PortHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ports/", h.GetPort)
	mux.HandleFunc("/ports/bulk", h.BulkUploadPorts)
}
