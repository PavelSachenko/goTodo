package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"newExp/internal/config"
	"newExp/internal/controller/http"
	"newExp/internal/repository"
	server2 "newExp/internal/server"
	"newExp/internal/usecase"
	"newExp/pkg/db/mysql"
)

type Claims struct {
	jwt.StandardClaims
	UserId uint64 `json:"userId"`
}

func main() {

	config, _ := config.Init("internal/config")
	db := mysql.NewMySqlConnection(
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Database,
	)
	repo := repository.NewSuperRepository(db)
	service := usecase.NewSuperService(repo)
	handler := http.NewHandler(service)
	server := server2.NewServer(config, handler.Init(config))
	err := server.Run()
	if err != nil {
		fmt.Println(err)
	}
	err = server.Stop()
}
