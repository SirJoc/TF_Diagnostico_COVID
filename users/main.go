// first one :: go get github.com/gofiber/fiber/v2
// second one :: go get -u gorm.io/gorm
// third one :: go get -u gorm.io/driver/mysql

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id uint
	FirstName string
	LastName string
	Username string
	Password string
}

func main() {
	// gorm
	db, err := gorm.Open(mysql.Open("root:password@tcp(127.0.0.1:3306)/users_ms"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(User{})


	app := fiber.New()
	app.Use(cors.New())


	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Users ms! ")
	})
	app.Listen(":8000")
}
