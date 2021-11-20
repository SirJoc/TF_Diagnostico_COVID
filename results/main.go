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
	DolorGarganta bool `json:"dolor_garganta"`
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
	ResultId uint `json:"result_id"`
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

	app.Get("/api/users/:id/results", func(c *fiber.Ctx) error {
		var results []Result
		db.Find(&results, "user_id = ?", c.Params("id"))
		for i, result := range results {
			response, err := http.Get(fmt.Sprintf("http://localhost:8002/api/results/%d/forms", result.Id))
			if err != nil {
				return err
			}

			var form Form
			json.NewDecoder(response.Body).Decode(&form)
			fmt.Println(form)
			results[i].Form = form
		}
		return c.JSON(results)
	})
	app.Get("/api/results", func(c *fiber.Ctx) error {
		var results []Result
		db.Find(&results)
		return c.JSON(results)
	})

	app.Post("/api/results", func(c *fiber.Ctx) error {
		var result Result
		if err := c.BodyParser(&result); err!= nil {
			return err
		}
		db.Create(&result)
		return c.JSON(result)
	})


	app.Listen(":8001")
}