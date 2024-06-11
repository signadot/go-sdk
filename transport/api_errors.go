package transport

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/signadot/go-sdk/models"
)

// FixAPIErrors is middleware that extracts the remote, server-provided API
// error message from an SDK error.
func FixAPIErrors(transport runtime.ClientTransport) runtime.ClientTransport {
	return apiErrorTransport{base: transport}
}

type apiErrorTransport struct {
	base runtime.ClientTransport
}

func (t apiErrorTransport) Submit(op *runtime.ClientOperation) (interface{}, error) {
	op.Reader = apiErrorReader{base: op.Reader}
	return t.base.Submit(op)
}

type apiErrorReader struct {
	base runtime.ClientResponseReader
}

func (r apiErrorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	code := response.Code()
	switch {
	case code >= 400 && code <= 599:

		body, err := io.ReadAll(response.Body())
		if err != nil {
			return nil, fmt.Errorf("can't read response body: %w", err)
		}
		apiErr := &APIError{}
		if err := json.Unmarshal(body, apiErr); err != nil {
			// If the response is not valid JSON, just return the raw body.
			return nil, fmt.Errorf("%v: %v", response.Message(), string(body))
		}
		apiErr.Code = int64(code)
		return nil, fmt.Errorf("%v: %w", response.Message(), apiErr)
	default:
		return r.base.ReadResponse(response, consumer)
	}
}

type APIError struct {
	models.ErrorResponse
}

func (e *APIError) Error() string {
	return e.ErrorResponse.Error
}
