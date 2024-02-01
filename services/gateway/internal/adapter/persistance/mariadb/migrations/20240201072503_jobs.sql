-- +goose Up
-- +goose StatementBegin
CREATE TABLE jobs (
    id CHAR(36) NOT NULL,
    target VARCHAR(255) NOT NULL,
    payload JSON NOT NULL,
    status VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE jobs;
-- +goose StatementEnd