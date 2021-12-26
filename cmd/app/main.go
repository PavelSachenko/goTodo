package main

import (
	"newExp/internal/config"
	"newExp/internal/controller/http"
	"newExp/internal/repository"
	server2 "newExp/internal/server"
	"newExp/internal/usecase"
	"newExp/pkg/db/mysql"
)

func main() {
	db := mysql.NewMySqlConnection("testing")
	config, _ := config.Init("internal/config")
	repo := repository.NewSuperRepository(db)
	service := usecase.NewSuperService(repo)
	handler := http.NewHandler(service)
	server := server2.NewServer(config, handler.Init(config))
	err := server.Run()
	if err != nil {
		panic(err)
	}
	err = server.Stop()
}
