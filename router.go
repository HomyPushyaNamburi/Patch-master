package server

import (
	"net/http"
	"patch-master/controllers"
	"strings"

	util "patch-master/util.go"

	"github.com/gin-gonic/gin"
)

// NewRouter :- Returns router mapppings
func NewRouter() *gin.Engine {
	router := gin.New()

	lctrl := new(controllers.LoginController)
	router.GET("/login", lctrl.Login)

	userRouter := router.Group("/user")
	userRouter.Use(AuthMiddleware())

	uctrl := new(controllers.UserController)
	userRouter.GET("", uctrl.Get)

	return router

}

// AuthMiddleware :-
func AuthMiddleware() func(*gin.Context) {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")
		slpit := strings.Split(accessToken, " ")
		err := util.ValidateJwtToken(slpit[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
		}
		c.Next()
	}
}
