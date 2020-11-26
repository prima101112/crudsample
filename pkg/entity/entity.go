package entity

import "time"

// Gopher is original gopher entity
type Gopher struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Company   string    `json:"company" db:"company"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
}

// Gophers many gopher
type Gophers []Gopher

// DefaultResponse is default response json
type DefaultResponse struct {
	Data    interface{} `json:"data"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
}
