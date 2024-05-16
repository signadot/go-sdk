package transport

import (
	"github.com/go-openapi/runtime"
)

type apiComposableTransport struct {
	base      runtime.ClientTransport
	artifacts runtime.ClientTransport
}

func (t apiComposableTransport) Submit(op *runtime.ClientOperation) (interface{}, error) {
	switch op.ID {
	case "download-job-attempt-artifact", "info-job-attempt-artifact", "list-job-attempt-artifacts", "upload-job-attempt-artifact":
		return t.artifacts.Submit(op)
	default:
		return t.base.Submit(op)
	}
}
