package port

import (
	"context"

	"go-kafka-demo/libs/domain"

	"github.com/google/uuid"
)

type JobRepository interface {
	CreateJob(ctx context.Context, job *domain.Job, tx func() error) error
	UpdateJobWithStatus(ctx context.Context, id uuid.UUID, status domain.JobStatus, tx func() error) error

	GetJob(ctx context.Context, id uuid.UUID) (*domain.Job, error)
}
