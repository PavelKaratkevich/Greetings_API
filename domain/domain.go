package domain

import "greetings/err"

// domain
type GreetingResponse struct {
	Slug string `json:"Name"`
	Video string `json:"Video"`
	Preview string `json:"Preview"`
}

// primary port
type GreetingRepository interface {
	GetGreeting(name string) ([]GreetingResponse, *err.Error)
}
