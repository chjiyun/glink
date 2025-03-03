package main

import (
	"Baiyuetribe/glink/service"
	"log"

	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		// DisableStartupMessage: true,
	})
	// app.Use(logger.New())
	app.Use(cors.New())
	app.Use(compress.New()) // 压缩静态资源未gzip或br
	app.Static("/", "web")  // 静态文件
	// app.Use("/", filesystem.New(filesystem.Config{
	// 	Root: packr.New("Assets Box", "/web"),
	// }))
	app.Get("api/*", service.ApiHandler) // 请求地址 http://127.0.0.1:3000/api/http://demo.com
	log.Fatal(app.Listen(":8006"))
}
