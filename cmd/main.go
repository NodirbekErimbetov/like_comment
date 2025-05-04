package main

import (
	"log"
	"minimedium/api"
	"minimedium/config"
	"minimedium/storage/postgres"

	"github.com/gin-gonic/gin"
)

func main() {

	var cfg = config.Load()

	pgStore, err := postgres.NewConnectionPostgres(&cfg)
	if err != nil {
		panic(err)
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	api.SetUpApi(r, &cfg, pgStore)
	log.Println("Listening "+cfg.ServiceHost+cfg.ServiceHTTPPort, "......")
	if err := r.Run(cfg.ServiceHost + cfg.ServiceHTTPPort); err != nil {
		panic("Listent and service panic:" + err.Error())

	}

}
