package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/gibran-abdillah/rest-jwt/models"
	"github.com/golang-jwt/jwt/v5"
)

func Get_user(user_id int) (models.User, error) {
	var user models.User

	if result := models.DB.First(&user, user_id); result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func GenerateJwtToken(username string) (string, error) {
	secret_key := os.Getenv("SECRET_KEY")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	claims["sub"] = 1

	token_string, err := token.SignedString([]byte(secret_key))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return token_string, nil
}

func CheckJWT(tokenString string) (interface{}, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header["sub"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return "", fmt.Errorf("cant validate your string")
}
