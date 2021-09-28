package app

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"greetings/handlers"
	"greetings/repository"
	"log"
	"net/http"
	"os"
	"time"
)

//  PORT_NAME=8000 ADDRESS_NAME=localhost go run .

func StartApp() {
	router := mux.NewRouter()

	db := ConnectDB()
	greetingRepositoryDb := repository.NewGreetingRepositoryDb(db)
	gs := handlers.NewGreetingService(greetingRepositoryDb)

	router.HandleFunc("/{name}", gs.GetGreeting).Methods(http.MethodGet)

	address := os.Getenv("ADDRESS_NAME")
	port := os.Getenv("PORT_NAME")

	log.Printf("Server is running at %s:%s", address, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func ConnectDB() *sqlx.DB {

	dataSource := fmt.Sprintf("root:I240959ko@tcp(localhost:3306)/iko_cards")
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil || client == nil {
		log.Fatal("Error while opening DB: ", err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}