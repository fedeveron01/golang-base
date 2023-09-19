package internal_jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	*jwt.StandardClaims
	TokenType  string
	EmployeeId float64
	SessionId  float64
	Role       string
}

func ParseToken(tokenString string) (Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("test"), nil
	})
	if err != nil {
		return Claims{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return Claims{
			Role:       claims["Role"].(interface{}).(string),
			EmployeeId: claims["EmployeeId"].(float64),
			SessionId:  claims["SessionId"].(float64),
		}, nil

	} else {
		fmt.Println(err)
	}
	return Claims{}, err
}
