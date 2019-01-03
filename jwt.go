package util

import (
	"fmt"
	"log"

	jwt "github.com/dgrijalva/jwt-go"

	"time"
)

const JWT_SECRET = "thisissdsdd"

func CreateJwtToken(user map[string]interface{}) (string, error) {
	hmacSampleSecret := []byte("JWT_SECRET")
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":    "Naveen",
		"user_id": "508",
		"exp":     time.Now().Add(time.Minute * 1).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(hmacSampleSecret)
}

// ValidateJwtToken :-
func ValidateJwtToken(tokenString string) error {
	hmacSampleSecret := []byte("JWT_SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Println(claims["name"])
		// will get the user id and checks the user in the db again and check the permission
		return nil
	} else {
		return err
	}

}
