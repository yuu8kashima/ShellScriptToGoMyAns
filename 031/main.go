package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type userType struct {
	Name string
}

// Fiberフレームワークを使用して、簡単なREST APIを作成してください。
// /usersエンドポイントでGETリクエストを処理し、ユーザーのリストをJSONで返すようにします。
// メモ: 本来はDBからselectするのだろうが、別問でやってるし、無しで。
func main() {
	users := []userType{}
	users = append(users, userType{"ABC-san"})

	app := fiber.New()
	app.Get("/users", func(c *fiber.Ctx) error {
		data, _ := json.Marshal(users)
		return c.SendString(string(data))
	})
	app.Listen(":3000")
}
