package domain

import (
	"time"

	"github.com/google/uuid"
)

type JobTarget string
type JobStatus string

const (
	JobTargetA JobTarget = "a"
	JobTargetB JobTarget = "b"
	JobTargetC JobTarget = "c"

	JobStatusHold      JobStatus = "hold"
	JobStatusPending   JobStatus = "pending"
	JobStatusCompleted JobStatus = "completed"
	JobStatusFailed    JobStatus = "failed"
	JobStatusCancelled JobStatus = "cancelled"
)

type Job struct {
	Id     uuid.UUID `json:"id"`
	Status JobStatus `json:"status"`

	Payload string `json:"payload"`

	Target    JobTarget `json:"target"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type JobOption func(*Job)

func WithStatus(status JobStatus) JobOption {
	return func(j *Job) {
		j.Status = status
	}
}

func NewJob(target JobTarget, opts ...JobOption) *Job {
	j := &Job{
		Id: uuid.Must(uuid.NewV7()),

		Status: JobStatusPending,
		Target: target,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	for _, opt := range opts {
		opt(j)
	}

	return j
}
