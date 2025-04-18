-- +goose Up
-- +goose StatementBegin
CREATE TABLE uploads
(
    uuid            TEXT NOT NULL PRIMARY KEY,
    signer_algo     TEXT NOT NULL,
    s3_key          TEXT NOT NULL,
    s3_urn          TEXT NOT NULL,
    byte_size       INTEGER,
    ttl             DATETIME NOT NULL,
    signature_hex   TEXT,
    created_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE uploads;
-- +goose StatementEnd
