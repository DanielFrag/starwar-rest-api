package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/DanielFrag/starwar-rest-api/infra"
	"github.com/DanielFrag/starwar-rest-api/router"
)

func main() {
	defer mainRecover()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		clean()
	}()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	infra.StartDB()
	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func clean() {
	defer func() {
		rec := recover()
		if rec != nil {
			log.Println(rec)
		}
		os.Exit(1)
	}()
	stopDb()
}

func stopDb() {
	infra.StopDB()
	fmt.Println("DB closed")
}

func mainRecover() {
	rec := recover()
	if rec != nil {
		log.Println(rec)
		clean()
	}
}
