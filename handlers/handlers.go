package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"greetings/domain"
	"greetings/dto"
	"log"
	"math/rand"
	"net/http"
)

type GreetingHandlers struct {
	repository domain.GreetingRepository
}

func (g *GreetingHandlers) GetGreeting(w http.ResponseWriter, r *http.Request) {
	var reply []dto.Response
	params := mux.Vars(r)
	name := params["name"]

	response, err := g.repository.GetGreeting(name)
	log.Println(response)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(err)
		return
	}

	q := rand.Intn(len(response)+1)

	k := dto.Response{
		Name:    name,
		Slug:    response[q].Slug,
		Preview: response[q].Preview,
		Video:   response[q].Video,
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