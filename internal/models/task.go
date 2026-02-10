package models

import "time"

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Done      bool   `json:"done"`
	CreatedAt string `json:"created_at,omitempty"`
}

func (t *Task) SetDefaults() {
	if t.CreatedAt == "" {
		t.CreatedAt = time.Now().Format(time.RFC3339)
	}
}
