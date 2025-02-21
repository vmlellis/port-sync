package processor

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"sync"

	"github.com/vmlellis/port-sync/src/internal/domain/contract"
	"github.com/vmlellis/port-sync/src/internal/domain/entity"
)

const (
	defaultBufferedChannelSize = 100
	defaultWorkersPoolSize     = 5
)

// ParallelProcessor implements the Processor interface.
type ParallelProcessor struct {
	service             contract.PortService
	bufferedChannelSize int
	workersPoolSize     int
	mu                  sync.Mutex
}

// parallelProcessorOpts defines the settings
type ParallelProcessorOpts struct {
	BufferedChannelSize int
	WorkersPoolSize     int
}

// NewParallelProcessor creates a new instance of ParallelProcessor.
func NewParallelProcessor(service contract.PortService, opts ParallelProcessorOpts) contract.ProcessorService {
	bufferedChannelSize := defaultBufferedChannelSize
	if opts.BufferedChannelSize > 0 {
		bufferedChannelSize = opts.BufferedChannelSize
	}

	workersPoolSize := defaultWorkersPoolSize
	if opts.WorkersPoolSize > 0 {
		workersPoolSize = opts.WorkersPoolSize
	}

	return &ParallelProcessor{
		service:             service,
		bufferedChannelSize: bufferedChannelSize,
		workersPoolSize:     workersPoolSize,
	}
}

func (p *ParallelProcessor) Process(rd io.Reader) error {

	// Use buffered reader for efficient file reads
	reader := bufio.NewReader(rd)
	decoder := json.NewDecoder(reader)

	// Read opening bracket
	if _, err := decoder.Token(); err != nil {
		return err
	}

	// Channel to pass parsed ports to worker goroutines
	portChannel := make(chan *entity.Port, p.bufferedChannelSize)
	var wg sync.WaitGroup

	// Start worker pool with concurrent workers
	for i := 0; i < p.workersPoolSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for port := range portChannel {
				p.service.SavePort(context.Background(), port) // Process and store each port
			}
		}()
	}

	for decoder.More() {
		var port entity.Port
		key, err := decoder.Token()
		if err != nil {
			return err
		}

		if id, ok := key.(string); ok {
			port.ID = id
		} else {
			return fmt.Errorf("unexpected key type: %T", key)
		}

		if err := decoder.Decode(&port); err != nil {
			return err
		}

		// Send the parsed port to worker pool
		portChannel <- &port
	}

	// Close channel after sending all ports
	close(portChannel)

	// Wait for all workers to finish
	wg.Wait()

	return nil
}
