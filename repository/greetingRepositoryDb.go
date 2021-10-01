package repository

import (
	"github.com/jmoiron/sqlx"
	"greetings/domain"
	"greetings/err"
	"log"
	"net/http"
)

type GreetingRepositoryDb struct {
	client *sqlx.DB
}

func (g GreetingRepositoryDb) GetGreetingByName(name string) ([]domain.GreetingResponse, *err.Error) {
	var error err.Error
	var greetingResponse []domain.GreetingResponse
	sqlRequest := "select cp.slug, cc.video, cc.preview from cardsmaker_card cc join cardsmaker_person cp on cc.name_id = cp.id where cp.name = ?"
	err := g.client.Select(&greetingResponse, sqlRequest, name)
		if err != nil {
				error.Message = "Error while retrieving data from database"
				error.Status = http.StatusInternalServerError
				log.Println("Error while retrieving database info: " + err.Error())
				return nil, &error
			}
	return greetingResponse, nil
}

func (g GreetingRepositoryDb) GetGreetingByAge(age string) ([]domain.GreetingResponse, *err.Error) {
	var error err.Error
	var greetingResponse []domain.GreetingResponse

	sqlRequest := "select name as age, cardname, ca.slug as description_eng, ca2.slug as number_of_years, video, preview from cardsmaker_agecard ca join cardsmaker_age ca2 on ca.name_id = ca2.id where ca2.slug = ?"

	err := g.client.Select(&greetingResponse, sqlRequest, age)
	if err != nil {
		error.Message = "Error while retrieving data from database"
		error.Status = http.StatusInternalServerError
		log.Println("error while retrieving database info: " + err.Error())
		return nil, &error
	}
	return greetingResponse, nil
}

func NewGreetingRepositoryDb(dbClient *sqlx.DB) GreetingRepositoryDb {
	return GreetingRepositoryDb{
		dbClient,
	}
}