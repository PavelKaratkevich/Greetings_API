package domain

import (
	"greetings/dto"
	"greetings/err"
)

// domain
type GreetingResponse struct {
	Name			string
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

func (d GreetingResponse) ToDto() dto.Response {
	return dto.Response{
		Name:            d.Name,
		Age:             d.Age,
		Cardname:        d.Cardname,
		Description_eng: d.Description_eng,
		Number_of_years: d.Number_of_years,
		Video:           d.Video,
		Preview:         d.Preview,
		Slug:            d.Slug,
	}
}