package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type ConnectionProvider interface {
	Connection() *sql.DB
	Run() error
	Stop()
}

type Config struct {
	DSN string
}

type DB struct {
	Config Config
	Conn   *sql.DB
}

func (d *DB) Run() error {
	if d.Conn != nil {
		return nil
	}
	if d.Config.DSN == "" {
		return fmt.Errorf("postgres dsn is required")
	}

	conn, err := sql.Open("postgres", d.Config.DSN)
	if err != nil {
		return err
	}
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxLifetime(30 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := conn.PingContext(ctx); err != nil {
		_ = conn.Close()
		return err
	}

	d.Conn = conn
	return nil
}

func (d *DB) Stop() {
	if d.Conn != nil {
		_ = d.Conn.Close()
	}
}

func (d *DB) Connection() *sql.DB {
	return d.Conn
}
