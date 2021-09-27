package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vitorbgouveia/go-rabbitmq/internal/usecases"
)

// HelloWord received and call usecase
func HelloWord(c *gin.Context) {
	var body []string
	c.ShouldBindJSON(&body)

	usecases.HelloWord(body)

	c.Status(200)
}
