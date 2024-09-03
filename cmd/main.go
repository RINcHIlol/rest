package main

import (
	restApi "github.com/RINcHIlol/rest.git"
	"github.com/RINcHIlol/rest.git/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(restApi.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
