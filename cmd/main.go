package main

import (
	todo "github.com/DiamondDmitriy/gopher-todo"
	"github.com/DiamondDmitriy/gopher-todo/internal/handler"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("server_port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
