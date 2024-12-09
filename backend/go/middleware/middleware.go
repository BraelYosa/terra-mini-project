package middleware

import (
	"backend/model"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

var JwtSkey = []byte("your_secret_key")

func GenerateJWT(user model.User, expiry time.Duration) (string, error) {

	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()

	// Create a new token object with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.User_id,
		"exp":     exp,
	})

	tokenString, err := token.SignedString(JwtSkey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CreateUser(newAuthor model.User) error {
	db := model.DB

	var existingUser model.User
	if err := db.Table("users").Where("User_name=?", newAuthor.UserSurname).First(&existingUser).Error; err == nil {
		return errors.New("author with the same name is already exists")
	}

	if err := db.Create(&newAuthor).Error; err != nil {
		return err
	}

	return nil
}

func GetuserByEmail(email string) (*model.User, error) {
	db := model.DB

	var user model.User
	if err := db.Table("users").Where("user_mail = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func FindUserByEmail(email string) (*model.User, error) {
	db := model.DB
	fmt.Println("Looking for user with email:", email) // Debug log

	var user model.User
	result := db.Where("user_mail = ?", email).First(&user)
	if result.Error != nil {
		fmt.Println("Error in FindUserByEmail:", result.Error)
		return nil, result.Error
	}

	return &user, nil
}

func CheckToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing token"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token format"})
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return JwtSkey, nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}

		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token expired"})
			}
		}

		// Add claims to context
		c.Set("user_id", claims["user_id"])
		return next(c)
	}
}
