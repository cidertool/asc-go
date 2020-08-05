package asc

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"sync"
)

// UploadOperation defines model for UploadOperation.
type UploadOperation struct {
	Length         *int                     `json:"length,omitempty"`
	Method         *string                  `json:"method,omitempty"`
	Offset         *int                     `json:"offset,omitempty"`
	RequestHeaders *[]UploadOperationHeader `json:"requestHeaders,omitempty"`
	URL            *string                  `json:"url,omitempty"`
}

// UploadOperationHeader defines model for UploadOperationHeader.
type UploadOperationHeader struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// Chunk returns the bytes in the file from the given offset and with the given length
func (op *UploadOperation) Chunk(f *os.File) (io.Reader, error) {
	_, err := f.Seek(int64(*op.Offset), 0)
	if err != nil {
		return nil, err
	}
	data := make([]byte, *op.Length)
	_, err = f.Read(data)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(data), nil
}

// Request creates a new http.Request instance from the given UploadOperation and buffer
func (op *UploadOperation) Request(data io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(*op.Method, *op.URL, data)
	if err != nil {
		return nil, err
	}
	for _, h := range *op.RequestHeaders {
		req.Header.Add(*h.Name, *h.Value)
	}
	return req, nil
}

// UploadMultipartFile takes a file path and some operations and concurrently uploads each part of the file to App Store Connect
func (c *Client) UploadMultipartFile(name string, operations []UploadOperation) error {
	// Open the file
	file, err := os.Open(name)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errs := make(chan error)

	for i, operation := range operations {
		chunk, err := operation.Chunk(file)
		if err != nil {
			errs <- err
			continue
		}
		wg.Add(1)
		go c.sendChunk(operations[i], chunk, errs, &wg)
	}

	go func() {
		wg.Wait()
		close(errs)
	}()

	// TODO: These errors ought to be concatenated in some way
	for err := range errs {
		return err
	}

	return nil
}

func (c *Client) sendChunk(op UploadOperation, chunk io.Reader, errs chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	req, err := op.Request(chunk)
	if err != nil {
		errs <- err
		return
	}
	_, err = c.do(req, nil)
	if err != nil {
		errs <- err
	}

}
