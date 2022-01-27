package main

import (
	// "Postresql/pkg/repository"
	"Postresql/pkg/handler"
	"Postresql/pkg/repository"
	"Postresql/pkg/services"
	"Postresql/server"
	"fmt"

	"github.com/spf13/viper"

	// "fmt"
	"log"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostSql(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     5432,
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBname:   viper.GetString("db.dbname"),
	})

	if err != nil {
		log.Fatalf("Error %s", err.Error())
	} else {
		fmt.Println("Succesfly connectde to database congruilations")
	}
	defer db.Close()

	repository := repository.NewRepository(db)
	service := services.NewService(repository)
	handler := handler.NewHandler(*service)
	server := new(server.Server)

	if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
