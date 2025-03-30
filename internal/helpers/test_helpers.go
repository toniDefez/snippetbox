package helpers

import (
	"database/sql"
	"fmt"
	"os/exec"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"snippetbox.tonidefez.net/internal/testdata"
)

func SetupTestDB(t *testing.T) (*sql.DB, error) {
	t.Helper()

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
		return nil, fmt.Errorf("failed to start container: %s\n%s", err, out)
	}

	dsn := "testuser:testpass@tcp(127.0.0.1:3307)/testdb?parseTime=true"

	var db *sql.DB
	var err error
	for i := 0; i < 15; i++ {
		db, err = sql.Open("mysql", dsn)
		if err == nil && db.Ping() == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %v", err)
	}

	t.Cleanup(func() {
		_ = db.Close()
		time.Sleep(500 * time.Millisecond)
		_ = exec.Command("docker", "rm", "-f", containerName).Run()
	})

	return db, nil
}

func LoadTestDataFromStruct(t *testing.T, db *sql.DB, data testdata.Seed) error {
	t.Helper()

	if data.Snippets != nil && len(data.Snippets) > 0 {
		_, err := db.Exec(`CREATE TABLE snippets (
			id INT AUTO_INCREMENT PRIMARY KEY,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			created DATETIME NOT NULL,
			expires DATETIME NOT NULL
		);`)
		if err != nil {
			return fmt.Errorf("error creating snippets table: %v", err)
		}

		for _, s := range data.Snippets {
			_, err := db.Exec(`
				INSERT INTO snippets (title, content, created, expires)
				VALUES (?, ?, ?, ?)
			`, s.Title, s.Content, s.Created, s.Expires)
			if err != nil {
				return fmt.Errorf("error inserting snippet: %v", err)
			}
		}
	}

	return nil
}
