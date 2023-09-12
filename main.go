package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aksentijevicd1/go-mysql-project/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	r := mux.NewRouter()
	routes.RegisterBooks(r)

	s := http.Server{
		Addr:         ":9090",
		Handler:      r,
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server %s:\n", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown!", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
