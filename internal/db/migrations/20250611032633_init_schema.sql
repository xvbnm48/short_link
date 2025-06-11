-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS links (
    id SERIAL PRIMARY KEY,
    short_code varchar(255) NOT NULL UNIQUE,
    original_url varchar(2048) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Jika kamu butuh trigger untuk update `updated_at`, lihat bagian bawah
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS clicks (
    id SERIAL PRIMARY KEY,
    link_id INT NOT NULL,
    ip_address VARCHAR(45) NOT NULL,
    user_agent VARCHAR(512) NOT NULL,
    clicked_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (link_id) REFERENCES links(id) ON DELETE CASCADE
);
-- SELECT 'up SQL query';
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS clicks;
DROP TABLE IF EXISTS links;
-- SELECT 'down SQL query';
-- +goose StatementEnd
