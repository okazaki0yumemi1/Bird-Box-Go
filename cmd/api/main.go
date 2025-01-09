package main

import (
	"fmt"
	"log"
	"net/http"

	"bird-box-go/api/router"
	"bird-box-go/config"
)

//  @title          Bird-Box-Go API
//  @version        1.0
//  @description    This is a work-in-progress version of a RESTful API with a CRUD functionality to control & access an app. Most things are taken from https://learning-cloud-native-go.github.io
//  @host       localhost:8080
//  @basePath   /v1

func main() {
	c := config.New()
	r := router.New()
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	log.Println("Starting server " + s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}
