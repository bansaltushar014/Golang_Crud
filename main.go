package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	route "github.com/bansaltushar014/GoLang_CRUD/routes"
	database "github.com/bansaltushar014/GoLang_CRUD/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	database.CreateConnection()
	run()
}

func run() error {
	mux := route.MakeMuxRouter()
	httpPort := os.Getenv("PORT")
	log.Println("HTTP Server Listening on port :", httpPort)
	s := &http.Server{
		Addr:           ":" + httpPort,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
