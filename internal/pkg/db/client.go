package db

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"tinvest-go/internal/pkg/logger"

	"database/sql"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

type IClient interface {
	GetDB() *sql.DB
	Close() error
}

type client struct {
	db *sql.DB
}

func NewClient(dsn string) (IClient, error) {
	if _, err := os.Stat(dsn); errors.Is(err, os.ErrNotExist) {
		_, filename, _, _ := runtime.Caller(0)
		fmt.Printf("Current test filename: %s\n", filename)
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		fmt.Println(exPath)
		return nil, fmt.Errorf("file db %s does not exist: %w", dsn, err)
	}
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create db client: %w", err)
	}
	return &client{
		db: db,
	}, nil
}

func (c *client) GetDB() *sql.DB {
	return c.db
}

func (c *client) Close() error {
	err := c.db.Close()
	if err != nil {
		return fmt.Errorf("failed to close db connection %w", err)
	}
	logger.Info(context.Background(), "db connection has been closed")
	return nil
}
