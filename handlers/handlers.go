package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"greetings/domain"
	"greetings/dto"
	"greetings/err"
	"greetings/utils"

	//err2 "greetings/err"
	_ "greetings/err"
	"log"
	"math/rand"
	"net/http"
)

type GreetingHandlers struct {
	repository domain.GreetingRepository
}

func (g *GreetingHandlers) GetGreetingByName(w http.ResponseWriter, r *http.Request) {
	var noNameFoundError err.Error // error for passing 'No name found' to the user
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
	reply := response[q].ToDto(name)
	utils.SendSuccess(w, reply)
}

func (g *GreetingHandlers) GetGreetingByAge(w http.ResponseWriter, r *http.Request) {
	var reply []dto.Response
	params := mux.Vars(r)
	age, _ := params["age"]
	log.Println(age)
	response, err := g.repository.GetGreetingByAge(age)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(err)
		return
	}

	q := rand.Intn(len(response)+1)
	k := dto.Response{
		Name:            "",
		Age:             age,
		Cardname:        response[q].Cardname,
		Description_eng: response[q].Description_eng,
		Number_of_years: response[q].Number_of_years,
		Video:           response[q].Video,
		Preview:         response[q].Preview,
		Slug:            response[q].Slug,
	}
	reply = append(reply, k)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reply)
}

func NewGreetingService(repository domain.GreetingRepository) GreetingHandlers {
	return GreetingHandlers{
		repository,
	}
}