package entity

// Gopher is original gopher entity
type Gopher struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Company string `json:"company"`
}

// Gophers many gopher
type Gophers []Gopher

// DefaultResponse is default response json
type DefaultResponse struct {
	Data    interface{} `json:"data"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
}
