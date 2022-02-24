package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/crissilvaeng/tedio/internal/core"
	"github.com/crissilvaeng/tedio/internal/misc"
	"github.com/gorilla/mux"
)

func main() {
	port := flag.String("port", misc.GetOrElseStr(os.Getenv("PORT"), "8080"), "port")
	apikey := flag.String("apikey", os.Getenv("API_KEY"), "api key")
	secret := flag.String("secret", os.Getenv("API_SECRET"), "api secret")
	flag.Parse()

	app, err := core.NewServer(core.Config{
		ApiKey:    *apikey,
		ApiSecret: *secret,
		Port:      *port,
	})

	if err != nil {
		log.Fatalf("error creating server: %v", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/games/{id}", app.Routes.GetGame).Methods(http.MethodGet)
	router.HandleFunc("/games", app.Routes.GetGames).Methods(http.MethodGet)
	router.HandleFunc("/games", app.Secure(app.Routes.PostGame)).Methods(http.MethodPost)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", *port),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		ErrorLog:     app.Logger,
	}

	app.Logger.Fatal(srv.ListenAndServe())
}
