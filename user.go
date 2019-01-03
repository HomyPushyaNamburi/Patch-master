package controllers

import (
	"github.com/gin-gonic/gin"
)

// UserController :-
type UserController struct {
}

// Get :-
func (us UserController) Get(c *gin.Context) {
	c.JSON(200, gin.H{"name": "Pushya"})
}
