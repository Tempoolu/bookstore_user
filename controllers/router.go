package controllers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	engine.GET("/user/:id", getUser)
}
