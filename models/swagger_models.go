package models

import "time"

// UserRegister represents the user registration request
type UserRegister struct {
	Username string `json:"username" example:"johndoe"`
	Email    string `json:"email" example:"john@example.com"`
	Password string `json:"password" example:"secretpass123"`
}

// UserResponse represents the user response
type UserResponse struct {
	ID       int64  `json:"id" example:"1"`
	Username string `json:"username" example:"johndoe"`
	Email    string `json:"email" example:"john@example.com"`
}

// LoginCredentials represents login request
type LoginCredentials struct {
	Username string `json:"username" example:"johndoe"`
	Password string `json:"password" example:"secretpass123"`
}

// LoginResponse represents login response
type LoginResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIs..."`
}

// ErrorResponse represents error response
type ErrorResponse struct {
	Error string `json:"error" example:"Invalid credentials"`
}

// ProjectResponse represents a project response
type ProjectResponse struct {
	ID          int64     `json:"id" example:"1"`
	Name        string    `json:"name" example:"My Project"`
	Description string    `json:"description" example:"A cool project"`
	UserID      int64     `json:"user_id" example:"1"`
	Visibility  string    `json:"visibility" example:"public"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
