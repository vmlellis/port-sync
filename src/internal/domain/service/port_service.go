package service

import (
	"github.com/vmlellis/port-sync/src/internal/domain/contract"
	"github.com/vmlellis/port-sync/src/internal/domain/entity"
)

// portService implements the PortService interface.
type portService struct {
	repo contract.PortRepository
}

// NewPortService creates a new instance of PortService and returns it as an interface.
func NewPortService(repo contract.PortRepository) contract.PortService {
	return &portService{repo: repo}
}

// SavePort inserts or updates a port record in the repository.
func (s *portService) SavePort(port *entity.Port) {
	s.repo.Save(port)
}

// GetPort retrieves a port by its unique identifier.
func (s *portService) GetPort(id string) (*entity.Port, bool) {
	return s.repo.Get(id)
}
