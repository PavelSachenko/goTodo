package main

import (
	"newExp/internal/config"
	"newExp/internal/controller/http"
	server2 "newExp/internal/server"
)

func main() {

	config, _ := config.Init("internal/config")
	handler := http.NewHandler()
	server := server2.NewServer(config, handler.Init(config))
	err := server.Run()
	if err != nil {
		panic(err)
	}
	err = server.Stop()
}
