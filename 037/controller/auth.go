package controller

import (
	"goshell-problem/037/crypt"
	"goshell-problem/037/database"
	"goshell-problem/037/model"

	"github.com/gofiber/fiber/v2"
)

type authRequest struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	var req authRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	user := model.User{
		PasswordHash: crypt.GeneratePassword(req.Password),
	}
	res := database.DB.Create(&user)
	if res.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": res.Error.Error(),
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "user created",
	})
}

func Login(c *fiber.Ctx) error {
	var req authRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	var user model.User
	res := database.DB.Where("id = ?", req.Id).First(&user)
	if res.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "user not found",
		})
	}
	if !crypt.ComparePassword(user.PasswordHash, req.Password) {
		return c.Status(400).JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	token, err := crypt.GenerateToken(user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"token": token,
	})
}
