package transport

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UploadPRGImageResult is the parsed response from a PRG image upload.
type UploadPRGImageResult struct {
	Digest string   `json:"digest"`
	Refs   []string `json:"refs,omitempty"`
	Size   int64    `json:"size"`
}

// UploadPRGImage streams an OCI archive body to the PRG image upload
// endpoint via the given transport. This bypasses the generated client
// method because swaggo doesn't generate a body parameter for
// application/x-tar consume types.
func UploadPRGImage(t runtime.ClientTransport, orgName, prgName, ref string, body io.Reader) (*UploadPRGImageResult, error) {
	op := &runtime.ClientOperation{
		ID:                 "upload-prg-image",
		Method:             "POST",
		PathPattern:        "/orgs/{orgName}/plan-runnergroups/{planRunnerGroupName}/images/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-tar"},
		Schemes:            []string{"https"},
		Params: &uploadPRGImageWriter{
			orgName: orgName,
			prgName: prgName,
			ref:     ref,
			body:    body,
		},
		Reader: &uploadPRGImageReader{},
	}
	result, err := t.Submit(op)
	if err != nil {
		return nil, err
	}
	resp, ok := result.(*UploadPRGImageResult)
	if !ok {
		return nil, fmt.Errorf("unexpected response type: %T", result)
	}
	return resp, nil
}

// uploadPRGImageWriter implements runtime.ClientRequestWriter.
type uploadPRGImageWriter struct {
	orgName string
	prgName string
	ref     string
	body    io.Reader
}

func (w *uploadPRGImageWriter) WriteToRequest(r runtime.ClientRequest, _ strfmt.Registry) error {
	if err := r.SetPathParam("orgName", w.orgName); err != nil {
		return err
	}
	if err := r.SetPathParam("planRunnerGroupName", w.prgName); err != nil {
		return err
	}
	if w.ref != "" {
		if err := r.SetQueryParam("ref", w.ref); err != nil {
			return err
		}
	}
	return r.SetBodyParam(w.body)
}

// uploadPRGImageReader implements runtime.ClientResponseReader.
type uploadPRGImageReader struct{}

func (rd *uploadPRGImageReader) ReadResponse(response runtime.ClientResponse, _ runtime.Consumer) (any, error) {
	code := response.Code()
	body, err := io.ReadAll(response.Body())
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}
	if code != http.StatusOK {
		return nil, fmt.Errorf("upload failed (HTTP %d): %s", code, string(body))
	}
	var result UploadPRGImageResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	return &result, nil
}
