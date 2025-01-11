package main

import (
	"goshell-problem/037/controller"
	"goshell-problem/037/database"
	"log"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// RESTful APIを作成し、JWTを使用した認証機能を実装してください。ユーザーはログイン情報を送信してトークンを取得し、
// そのトークンを使用して保護されたエンドポイントにアクセスできるようにします。
// 参考資料:
//   - https://medium.com/code-beyond/go-fiber-jwt-auth-eab51a7e2129
//   - 上記のコードサンプル https://github.com/zahidcakici/go-fiber-jwt
//   - https://qiita.com/koji0705/items/49172d713e13fa554ba7
//   - https://zenn.dev/joo_hashi/articles/83ebe09cb2fc59
func main() {
	readEnv()
	database.Connect()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")
	hello := api.Group("/hello")
	hello.Use(JWTProtected)
	hello.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Auth API!")
	})

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", controller.Login)
	auth.Post("/register", controller.Register)

	log.Fatal(app.Listen("localhost:3000"))
}

func readEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

func JWTProtected(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ContextKey: "jwt",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Return status 401 and failed authentication error.
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		},
	})(c)
}
