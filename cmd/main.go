package main

import (
	restApi "github.com/RINcHIlol/rest.git"
	"github.com/RINcHIlol/rest.git/pkg/handler"
	"github.com/RINcHIlol/rest.git/pkg/repository"
	"github.com/RINcHIlol/rest.git/pkg/service"
	_ "github.com/golang-migrate/migrate"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(restApi.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
