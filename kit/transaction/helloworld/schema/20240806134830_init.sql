-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS greater (
    hello VARCHAR(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS account;
-- +goose StatementEnd
