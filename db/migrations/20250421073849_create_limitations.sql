-- +goose Up
-- +goose StatementBegin
CREATE TABLE limitations(
    date                DATE PRIMARY KEY,
    written_bytes       INTEGER,
    written_counter     INTEGER,
    read_bytes          INTEGER,
    read_counter        INTEGER,
    created_at          DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE limitations;
-- +goose StatementEnd
