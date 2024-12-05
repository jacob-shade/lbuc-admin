package handler

import (
	"github.com/jacobshade/lbuc-admin/server/auth"

	"github.com/gofiber/fiber/v2"
)

// Auth fiber handler
// Redirects to Google login URL
func Auth(c *fiber.Ctx) error {
	path := auth.ConfigGoogle()
	url := path.AuthCodeURL("state")
	return c.Redirect(url)
}

// Callback to receive google's response
// Handles Google OAuth2 callback and returns user's email
func Callback(c *fiber.Ctx) error {
	token, error := auth.ConfigGoogle().Exchange(c.Context(), c.FormValue("code"))
	if error != nil {
		panic(error)
	}
	email := auth.GetEmail(token.AccessToken)
	return c.Status(200).JSON(fiber.Map{"email": email, "login": true})
}