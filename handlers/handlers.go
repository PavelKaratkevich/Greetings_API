package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"greetings/domain"
	"net/http"
)

type GreetingHandlers struct {
	repository domain.GreetingRepository
}

func (g *GreetingHandlers) GetGreeting(w http.ResponseWriter, r *http.Request) {
params := mux.Vars(r)
name := params["slug"]
//log.Println(name)

response, err := g.repository.GetGreeting(name)
if err != nil {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(err.Status)
	json.NewEncoder(w).Encode(err)
} else {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
}

func NewGreetingService(repository domain.GreetingRepository) GreetingHandlers {
	return GreetingHandlers{
		repository,
	}
}