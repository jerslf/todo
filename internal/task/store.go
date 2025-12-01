package task

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	Conn *sql.DB
}

// NewDB opens the SQLite file and ensures the table exists
func NewDB(path string) (*DB, error) {
	Conn, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// Create table if it doesn't exist
	_, err = Conn.Exec(`
        CREATE TABLE IF NOT EXISTS tasks (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT NOT NULL,
            done BOOLEAN NOT NULL DEFAULT 0,
            time_created DATETIME NOT NULL,
            time_done DATETIME
        )
    `)
	if err != nil {
		return nil, err
	}

	return &DB{Conn: Conn}, nil
}

// Close connection
func (db *DB) Close() error {
	return db.Conn.Close()
}

func (db *DB) Add(title string) (Task, error) {
	// Check empty string
	if strings.TrimSpace(title) == "" {
		return Task{}, errors.New("title cannot be empty")
	}
	// Check length
	if len(title) > 100 {
		return Task{}, errors.New("title is too long")
	}
	// Normalise
	title = strings.TrimSpace(title)

	now := time.Now()
	res, err := db.Conn.Exec("INSERT INTO tasks (title, done, time_created) VALUES (?, 0, ?)", title, now)
	if err != nil {
		return Task{}, err
	}

	id, _ := res.LastInsertId()
	return Task{ID: int(id), Title: title, Done: false, TimeCreated: now}, nil
}

func (db *DB) List(all bool) ([]Task, error) {
	query := "SELECT id, title, done, time_created, time_done FROM tasks"
	if !all {
		query += " WHERE done = 0"
	}

	rows, err := db.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		var timeDone sql.NullTime

		if err := rows.Scan(&t.ID, &t.Title, &t.Done, &t.TimeCreated, &timeDone); err != nil {
			return nil, err
		}
		if timeDone.Valid {
			t.TimeDone = &timeDone.Time
		}

		tasks = append(tasks, t)
	}

	return tasks, nil
}

// Mark as done / undone
func (db *DB) SetDone(id int, done bool) error {
	var timeDone interface{}
	if done {
		timeDone = time.Now()
	} else {
		timeDone = nil
	}
	res, err := db.Conn.Exec(
		"UPDATE tasks SET done=?, time_done=? WHERE id=?",
		done, timeDone, id,
	)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("task not found")
	}
	return nil
}

// Delete task
func (db *DB) Delete(id int) error {
	res, err := db.Conn.Exec("DELETE FROM tasks WHERE id=?", id)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("task not found")
	}
	return nil
}
