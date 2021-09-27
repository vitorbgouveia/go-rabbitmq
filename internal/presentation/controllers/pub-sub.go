package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vitorbgouveia/go-rabbitmq/internal/usecases"
)

// PubSub received and call usecase
func PubSub(c *gin.Context) {
	var body []string
	c.ShouldBindJSON(&body)

	usecases.PubSub(body)

	c.Status(200)
}
