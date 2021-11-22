// first one :: go get github.com/gofiber/fiber/v2
// second one :: go get -u gorm.io/gorm
// third one :: go get -u gorm.io/driver/mysql

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type User struct {
	Id uint `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Results []Result `json:"results" gorm:"-" default: "[]"`
}

type Result struct {
	Id uint `json:"id"`
	CreatedAt string `json:"created_at"`
	Diagnostic string `json:"diagnostic"`
	UserId uint `json:"user_id"`
	Form Form `json:"form"`
}

type Form struct{
	Id uint `json:"id"`
	Tos bool `json:"tos"`
	Cafelea bool `json:"cafelea"`
	CongNasal bool `json:"cong_nasal"`
	DifRespiratoria bool `json:"dif_respiratoria"`
	DolorGargante bool `json:"dolor_gargante"`
	Fiebre bool `json:"fiebre"`
	Diarrea bool `json:"diarrea"`
	Nauseas bool `json:"nauseas"`
	AnosmiaHiposmia bool `json:"anosmia_hiposmia"`
	DolorAbdominal bool `json:"dolor_abdominal"`
	DolorArticulaciones bool `json:"dolor_articulaciones"`
	DolorMuscular bool `json:"dolor_muscular"`
	DolorPecho bool `json:"dolor_pecho"`
	Otros bool `json:"otros"`
	Semanas uint `json:"semanas"`
}


func main() {
	// gorm
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/users_ms"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(User{})


	app := fiber.New()
	app.Use(cors.New())

	app.Get("/api/users", func(c *fiber.Ctx) error {
		var users []User

		db.Find(&users)
		for i, user := range users {
			response, err := http.Get(fmt.Sprintf("http://localhost:8001/api/users/%d/results/", user.Id))
			if err != nil {
				return err
			}
			var results []Result
			json.NewDecoder(response.Body).Decode(&results)
			users[i].Results = results
		}
		return c.JSON(users)
	})

	app.Post("/api/users", func(c *fiber.Ctx) error {
		var user User
		if err := c.BodyParser(&user); err != nil {
			return err
		}
		db.Create(&user)
		return c.JSON(user)
	})

	app.Get("/api/users/:username", func(c *fiber.Ctx) error {
		var user User
		db.Find(&user, "username = ?", c.Params("username"))

		return c.JSON(user.Password)
	})

	app.Listen(":8000")
}
