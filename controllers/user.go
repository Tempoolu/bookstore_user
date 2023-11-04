package controllers

import (
	"net/http"
    "encoding/json"

	"github.com/Tempoolu/bookstore_user/models"
	"github.com/Tempoolu/bookstore_user/clients"
	"github.com/gin-gonic/gin"
)

func getUser(c *gin.Context) {
	var user models.User
	user.Get(c.Param("id"))
    var result models.BookResult
    resp := clients.Rpc("book", "GET", "/book")
    _ = json.Unmarshal(resp, &result)
    user.Books = result.Book
	c.JSON(http.StatusOK, gin.H{
		"result": "success",
		"user":   user,
	})
}
