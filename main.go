package main

import (
	"github.com/gin-gonic/gin"
	"github.com/opn-ooo/gin-boilerplate/app/controller/shorten"
	"github.com/opn-ooo/gin-boilerplate/config"
	"github.com/opn-ooo/gin-boilerplate/config/database"
	"log"
)

func main() {
	// init starter
	goDotENV := config.GetGoDotENV()
	database.CreateDatabaseConnection(goDotENV.PostgresConfig)

	// start gin web server
	gin.SetMode(goDotENV.GinMode)
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	shortenR := router.Group("/s")
	{
		shortenR.POST("/encode", shorten.PostShorten)
		shortenR.GET("/:url", shorten.GetShortenRedirect)
	}

	log.Printf("\n\n PORT: %s \n ENV: %s \n", goDotENV.Port, goDotENV.GinMode)
	if err := router.Run(goDotENV.Port); err != nil {
		log.Fatal("Gin run error", err)
	}
}
