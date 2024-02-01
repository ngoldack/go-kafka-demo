package mariadb

import (
	"context"
	"errors"
	"go-kafka-demo/libs/domain"
	"go-kafka-demo/services/gateway/internal/app/port"

	"github.com/google/uuid"
)

// force implementation of port.JobRepository
var _ port.JobRepository = (*DB)(nil)

// CreateJob implements port.JobRepository.
func (d *DB) CreateJob(ctx context.Context, job *domain.Job, txFunc func() error) error {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	sql := `--sql
	INSERT INTO jobs (id, target, payload, status, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err = tx.ExecContext(ctx, sql, job.Id, job.Target, job.Payload, job.Status, job.CreatedAt, job.UpdatedAt)
	if err != nil {
		return err
	}

	err = txFunc()
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// GetJob implements port.JobRepository.
func (d *DB) GetJob(ctx context.Context, id uuid.UUID) (*domain.Job, error) {
	panic("unimplemented")
}

// UpdateJob implements port.JobRepository.
func (d *DB) UpdateJobWithStatus(ctx context.Context, id uuid.UUID, status domain.JobStatus, txFunc func() error) error {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	sql := `--sql
	UPDATE JobService
	SET status = ?
	WHERE id = ?
	`

	_, err = tx.ExecContext(ctx, sql, status, id)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return errors.Join(rbErr, err)
		}
		return err
	}

	err = txFunc()
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
