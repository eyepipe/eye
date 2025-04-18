package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/eyepipe/eye/internal/pkg/store"
	"github.com/pressly/goose/v3"
)

func SetupGooseDatabases(ctx context.Context, server *Server) ([]*sql.DB, error) {
	goose.SetVerbose(true)
	goose.WithAllowMissing()
	err := goose.SetDialect("sqlite3")
	if err != nil {
		return nil, fmt.Errorf("failed to goose.SetDialect: %w", err)
	}

	databases := make([]*sql.DB, 0, len(server.config.DDShardFiles))
	for _, filename := range server.config.DDShardFiles {
		var db *sql.DB
		db, err = store.Connect(filename)
		switch {
		case err != nil:
			return nil, fmt.Errorf("failed to connect to %s: %w", filename, err)
		default:
			databases = append(databases, db)
		}
	}

	return databases, nil
}
