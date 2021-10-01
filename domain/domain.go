package domain

import (
	"greetings/dto"
	"greetings/err"
)

// domain
type GreetingResponse struct {
	Age 			string
	Cardname 		string
	Description_eng string
	Number_of_years string
	Video 			string
	Preview 		string
	Description_rus string
	Slug 			string
}

// primary port
type GreetingRepository interface {
	GetGreetingByName(name string) ([]GreetingResponse, *err.Error)
	GetGreetingByAge(age string) ([]GreetingResponse, *err.Error)
}


func (d GreetingResponse) ToDto(name string) dto.Response {
	return dto.Response{
		Name:    name,
		Slug:    d.Slug,
		Preview: d.Preview,
		Video:   d.Video,
	}
}