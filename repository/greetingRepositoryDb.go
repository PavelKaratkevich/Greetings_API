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

func (g GreetingRepositoryDb) GetGreeting(name string) ([]domain.GreetingResponse, *err.Error) {
	var error err.Error
	sqlRequest := "select cp.slug, cc.video, cc.preview from cardsmaker_card cc join cardsmaker_person cp on cc.name_id = cp.id where cp.slug = ?"

var greetingResponse []domain.GreetingResponse

err := g.client.Select(&greetingResponse, sqlRequest, name)
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