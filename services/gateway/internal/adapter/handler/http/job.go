package http

import (
	"net/http"

	"go-kafka-demo/services/gateway/api"

	"github.com/google/uuid"
)

type JobHandler struct {
}

// CreateJob implements api.ServerInterface.
func (*JobHandler) CreateJob(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetJob implements api.ServerInterface.
func (*JobHandler) GetJob(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	panic("unimplemented")
}

// UpdateJobStatus implements api.ServerInterface.
func (*JobHandler) UpdateJobStatus(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	panic("unimplemented")
}

var _ api.ServerInterface = (*JobHandler)(nil)
