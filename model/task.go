package model

type Task struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
	Search    string `json:"search"`
}
