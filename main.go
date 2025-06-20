package main

import (
	"fmt"
	"go-21/database"
	"go-21/handler"
	"go-21/repository"
	"go-21/router"
	"go-21/service"
	"go-21/utils"
	"log"
	"net/http"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	// init logger
	logger, err := utils.InitLogger("./logs/app-", true)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db, logger)
	service := service.NewService(repo, logger)
	handler := handler.NewHandler(service)

	r := router.NewRouter(handler)

	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("error server")
	}
}
