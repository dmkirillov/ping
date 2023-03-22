package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	log.Info().Msg("Starting server")
	err := r.Run(":8080")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
