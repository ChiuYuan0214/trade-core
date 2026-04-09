package postgres

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Migrator interface {
	Run() error
	Stop()
}

type MigrationRunner struct {
	Database      ConnectionProvider
	MigrationsDir string
}

func (m *MigrationRunner) Run() error {
	if m.Database == nil || m.Database.Connection() == nil {
		return fmt.Errorf("migration runner db is nil")
	}
	if m.MigrationsDir == "" {
		return fmt.Errorf("migrations dir is required")
	}

	entries, err := os.ReadDir(m.MigrationsDir)
	if err != nil {
		return err
	}

	files := make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".sql") {
			continue
		}
		files = append(files, filepath.Join(m.MigrationsDir, entry.Name()))
	}
	sort.Strings(files)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	for _, file := range files {
		sqlBytes, err := os.ReadFile(file)
		if err != nil {
			return err
		}
		if len(strings.TrimSpace(string(sqlBytes))) == 0 {
			continue
		}
		if _, err := m.Database.Connection().ExecContext(ctx, string(sqlBytes)); err != nil {
			return fmt.Errorf("apply migration %s: %w", file, err)
		}
	}

	return nil
}

func (m *MigrationRunner) Stop() {}
