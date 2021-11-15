package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Result struct {
	Id uint
	CreateAt string
	Diagnostic string
	UserId uint
}


func main() {
	// gorm
	db, err := gorm.Open(mysql.Open("root:password@tcp(127.0.0.1:3306)/results_ms"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(Result{})


	app := fiber.New()
	app.Use(cors.New())


	app.Get("/", func(c *fiber.Ctx) error {
	return c.SendString("Results ms")
	})
	app.Listen(":8001")
}