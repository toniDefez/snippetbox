package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type TestSnippet struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Expires int    `json:"expires"`
}

func startMySQLContainer(t *testing.T) {
	containerName := fmt.Sprintf("mysql-test-%d", time.Now().UnixNano())

	cmd := exec.Command(
		"docker", "run", "-d", "--rm",
		"--name", containerName,
		"-e", "MYSQL_ROOT_PASSWORD=secret",
		"-e", "MYSQL_DATABASE=testdb",
		"-e", "MYSQL_USER=testuser",
		"-e", "MYSQL_PASSWORD=testpass",
		"-p", "3307:3306",
		"mysql:8.0",
	)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to start container: %s\n%s", err, out)
	}

	// Cleanup automático al finalizar el test
	t.Cleanup(func() {
		exec.Command("docker", "rm", "-f", containerName).Run()
	})
}

func waitForMySQLReady(t *testing.T, dsn string) *sql.DB {
	var db *sql.DB
	var err error

	for i := 0; i < 15; i++ {
		db, err = sql.Open("mysql", dsn)
		if err == nil && db.Ping() == nil {
			return db
		}
		time.Sleep(2 * time.Second)
	}
	t.Fatalf("MySQL not ready after retries: %v", err)
	return nil
}

func loadTestDataFromJSON(t *testing.T, db *sql.DB, path string) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		t.Fatalf("cannot resolve absolute path: %v", err)
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		t.Fatalf("error reading testdata file: %v", err)
	}

	var snippets []TestSnippet
	if err := json.Unmarshal(data, &snippets); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	for _, s := range snippets {
		_, err := db.Exec(`
			INSERT INTO snippets (title, content, created, expires)
			VALUES (?, ?, ?, DATE_ADD(?, INTERVAL ? DAY))
		`, s.Title, s.Content, time.Now(), time.Now(), s.Expires)
		if err != nil {
			t.Fatalf("error inserting test snippet: %v", err)
		}
	}
}
