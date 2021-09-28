package domain

import "greetings/err"

// domain
type GreetingResponse struct {
	Slug string `db:"slug"`
	Video string
	Preview string
}

// primary port
type GreetingRepository interface {
	GetGreeting(name string) ([]GreetingResponse, *err.Error)
}
