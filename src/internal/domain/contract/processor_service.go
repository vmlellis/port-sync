package contract

import "io"

// ProcessorService interface
type ProcessorService interface {
	// Process a reader
	Process(rd io.Reader) error
}
