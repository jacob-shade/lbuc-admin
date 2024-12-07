package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/jacobshade/lbuc-admin/server/config"
	"github.com/jacobshade/lbuc-admin/server/database"
	"github.com/jacobshade/lbuc-admin/server/model"

	"github.com/gofiber/fiber/v2"
)

// Redirects to Google login URL
func GoogleLogin(c *fiber.Ctx) error {
	path := config.GoogleConfig()
	url := path.AuthCodeURL("randomstate")
	return c.Redirect(url)
}

// Callback to receive google's response
// Handles Google OAuth2 callback and returns user's email
// func GoogleCallback(c *fiber.Ctx) error {
// 	token, error := config.GoogleConfig().Exchange(c.Context(), c.FormValue("code"))
// 	if error != nil {
// 		panic(error)
// 	}
// 	email := auth.GetEmail(token.AccessToken)
// 	return c.Status(200).JSON(fiber.Map{"email": email, "login": true})
// }

func GoogleCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != "randomstate" {
		return c.SendString("States don't Match!!")
	}

	code := c.Query("code")
	googlecon := config.GoogleConfig()

	token, err := googlecon.Exchange(context.Background(), code)
	if err != nil {
		return c.SendString("Code-Token Exchange Failed")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return c.SendString("User Data Fetch Failed")
	}

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.SendString("JSON Parsing Failed")
	}

	var googleUser model.GoogleResponse
	if err := json.Unmarshal(userData, &googleUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse user data"})
	}

	// Create session
	sess, err := database.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create session"})
	}

	// Set session values
	sess.Set("email", googleUser.Email)
	sess.Set("user_id", googleUser.ID)

	// Save session
	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save session"})
	}

	// return c.Status(fiber.StatusOK).JSON(fiber.Map{
	// 	"email": googleUser.Email,
	// 	"login": true,
	// })
	return c.Redirect(os.Getenv("CLIENT_URL"), fiber.StatusTemporaryRedirect)
}

func GetSession(c *fiber.Ctx) error {
	sess, err := database.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"authenticated": false,
			"error":         "No session found",
		})
	}

	email := sess.Get("email")
	userID := sess.Get("user_id")

	if email == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"authenticated": false,
			"error":         "Not authenticated",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"authenticated": true,
		"email":         email,
		"user_id":       userID,
	})
}

func Signout(c *fiber.Ctx) error {
	sess, err := database.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "No session found"})
	}

	if err := sess.Destroy(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to destroy session"})
	}

	return c.SendStatus(fiber.StatusOK)
}
