package http

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/vmlellis/port-sync/src/internal/domain/contract"
	"github.com/vmlellis/port-sync/src/internal/domain/entity"
)

// PortHandler provides HTTP endpoints to interact with port data.
type PortHandler struct {
	service contract.PortService
}

// NewPortHandler creates a new instance of PortHandler.
func NewPortHandler(service contract.PortService) *PortHandler {
	return &PortHandler{service: service}
}

// GetPort retrieves a port by its ID.
func (h *PortHandler) GetPort(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 || pathParts[1] != "ports" {
		http.Error(w, "Invalid request path", http.StatusBadRequest)
		return
	}
	id := pathParts[2]

	port, found := h.service.GetPort(id)
	if !found {
		http.Error(w, "Port not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(port)
}

// BulkUploadPorts handles bulk upload of ports from a JSON file.
func (h *PortHandler) BulkUploadPorts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var ports map[string]entity.Port
	if err := json.Unmarshal(body, &ports); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	for id, port := range ports {
		port.ID = id
		h.service.SavePort(&port)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Ports uploaded successfully"))
}

// RegisterRoutes sets up the HTTP routes for port operations.
func (h *PortHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ports/", h.GetPort)
	mux.HandleFunc("/ports/bulk", h.BulkUploadPorts)
}
