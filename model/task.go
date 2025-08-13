package model

import "time"

type Task struct {
	Name        string    `json:"name"`
	Completed   bool      `json:"completed"`
	Search      string    `json:"search"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	AccessesAt  time.Time `json:"accessed_at"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
}
