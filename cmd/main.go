package main

import (
	restApi "github.com/RINcHIlol/rest.git"
	"github.com/RINcHIlol/rest.git/pkg/handler"
	"github.com/RINcHIlol/rest.git/pkg/repository"
	"github.com/RINcHIlol/rest.git/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(restApi.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
