package controllers

import (
	"backend/helpers"
	"backend/middleware"
	"backend/model"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c echo.Context) error {
	var user model.User

	if err := c.Bind(&user); err != nil {
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Error_Bind_user")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.UserPass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.UserPass = string(hashedPass)

	if err := middleware.CreateUser(user); err != nil {
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Error_creating_user")
	}

	token, err := middleware.GenerateJWT(user, 24)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, map[string]any{
		"Token": token,
	})

}

func Login(c echo.Context) error {

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Println("Error reading raw request body:", err)
		return helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid request")
	}
	fmt.Println("Raw request body:", string(body))

	// Reset the body for further use by Bind()
	c.Request().Body = io.NopCloser(bytes.NewBuffer(body))

	var userLogin model.User
	if err := c.Bind(&userLogin); err != nil {
		fmt.Println("Error binding login request:", err)
		return helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid request")
	}
	fmt.Printf("Parsed login data: UserEmail = %s, UserPass = %s\n", userLogin.UserEmail, userLogin.UserPass)

	user, err := middleware.GetuserByEmail(userLogin.UserEmail)
	if err != nil {
		fmt.Println("Error retrieving user:", err)
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Invalid username or password")
	}

	// Compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.UserPass), []byte(userLogin.UserPass)); err != nil {
		fmt.Println("Password mismatch")
		return helpers.ErrorResponse(c, http.StatusUnauthorized, "Invalid username or password")
	}

	token, err := middleware.GenerateJWT(*user, 24*time.Hour)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token")
	}

	expTime := time.Now().Add(24 * time.Hour)
	fmt.Println("Login successful. Token generated.")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token":  token,
		"Exp_at": expTime.Format(time.RFC3339),
	})
}

func Protected(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "You are authorized!",
	})
}

func GetData(c echo.Context) error {
	// Mock data for demonstration
	data := []map[string]interface{}{
		{"id": 1, "name": "Item 1"},
		{"id": 2, "name": "Item 2"},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}

func SubmitData(c echo.Context) error {
	var formData model.FormData

	if err := c.Bind(&formData); err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid input")
	}

	// Logic to save formData to database (mocked here)
	// Example: err := model.DB.Create(&formData).Error
	// Skipping actual DB save for brevity

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Data submitted successfully!",
	})
}
