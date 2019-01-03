package controllers

import (
	"patch-master/util"

	"github.com/gin-gonic/gin"
	util "patch-master/util.go"
)

//LoginController ...
type LoginController struct {
}

// func CreateJwtToken() (string, error) {
// 	hmacSampleSecret := []byte("thisissdsdd")
// 	// Create a new token object, specifying signing method and the claims
// 	// you would like it to contain.
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"name":    "Pushya",
// 		"user_id": "508",
// 		"exp":     time.Now().Add(time.Minute * 1).Unix(),
// 	})

// 	// Sign and get the complete encoded token as a string using the secret
// 	return token.SignedString(hmacSampleSecret)
// }

func (ctrl LoginController) Login(c *gin.Context) {
	// Login logic :- TODO
	//Here user comes after scuccessfull login - USER INFO will come
	// user_id = uniq

	// if user not found / pwd is incorrect will sent 401 resp

	// if user authneticated

	user := map[string]interface{}{"user_id": 101}

	token, _ := util.CreateJwtToken(user)
	c.JSON(200, gin.H{"token": token})
}
