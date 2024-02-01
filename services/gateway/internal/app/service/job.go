package service

import (
	"context"
	"fmt"

	"go-kafka-demo/libs/domain"
	"go-kafka-demo/services/gateway/internal/app/port"
)

type JobService struct {
	jobRepository port.JobRepository
	workerTargets map[domain.JobTarget]port.Publisher
}

func NewJobService(workerTargets map[domain.JobTarget]port.Publisher, jr port.JobRepository) *JobService {
	return &JobService{
		jobRepository: jr,
		workerTargets: workerTargets,
	}
}

func (s *JobService) ScheduleJob(ctx context.Context, job *domain.Job) error {
	var tx = func() error {
		var err error
		switch job.Target {
		case domain.JobTargetA:
			err = s.workerTargets[domain.JobTargetA].Publish(ctx, job)
		case domain.JobTargetB:
			err = s.workerTargets[domain.JobTargetB].Publish(ctx, job)
		default:
			err = fmt.Errorf("unknown job target: %s", job.Target)
		}
		if err != nil {
			return fmt.Errorf("failed to publish job: %w", err)
		}

		return nil
	}

	err := s.jobRepository.CreateJob(ctx, job, tx)
	if err != nil {
		return fmt.Errorf("failed to create job: %w", err)
	}

	return nil
}
