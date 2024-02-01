package port

import (
	"context"

	"go-kafka-demo/libs/domain"
)

type Publisher interface {
	Publish(ctx context.Context, job *domain.Job) error
}
