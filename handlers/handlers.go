package handlers

import (
	"github.com/gorilla/mux"
	"greetings/domain"
	"greetings/err"
	_ "greetings/err"
	"greetings/utils"
	"math/rand"
	"net/http"
)

type GreetingHandlers struct {
	repository domain.GreetingRepository
}

func (g *GreetingHandlers) GetGreetingByName(w http.ResponseWriter, r *http.Request) {
	var noNameFoundError err.Error
	params := mux.Vars(r)
	name := params["name"]
	response, err := g.repository.GetGreetingByName(name)
		if err != nil {
			utils.SendError(w, err.Status, *err)
			return
		} else if response == nil {
			noNameFoundError.Message = "No name found"
			noNameFoundError.Status = http.StatusNotFound
			utils.SendError(w, noNameFoundError.Status, noNameFoundError)
			return
		}
	q := rand.Intn(len(response) + 1)
	reply := response[q].ToDto()
	utils.SendSuccess(w, reply)
}

func (g *GreetingHandlers) GetGreetingByAge(w http.ResponseWriter, r *http.Request) {
	var noAgeFoundError err.Error
	params := mux.Vars(r)
	age, _ := params["age"]
	response, err := g.repository.GetGreetingByAge(age)
	if err != nil {
		utils.SendError(w, err.Status, *err)
		return
	} else if response == nil {
		noAgeFoundError.Message = "No age found"
		noAgeFoundError.Status = http.StatusNotFound
		utils.SendError(w, noAgeFoundError.Status, noAgeFoundError)
		return
	}
	q := rand.Intn(len(response)+1)
	reply := response[q].ToDto()
	utils.SendSuccess(w, reply)
}

func NewGreetingService(repository domain.GreetingRepository) GreetingHandlers {
	return GreetingHandlers{
		repository,
	}
}