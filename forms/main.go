package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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
	ResultId uint `json:"result_id"`
}

func main() {
	// gorm
	db, err := gorm.Open(mysql.Open("root:password@tcp(127.0.0.1:3306)/forms_ms"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(Form{})

	app := fiber.New()
	app.Use(cors.New())


	app.Post("/api/users/:id/results/forms", func(c *fiber.Ctx) error {
		var form Form
		if err := c.BodyParser(&form); err != nil {
			return err
		}
		db.Create(&form)
		return c.JSON(form)
	})

	app.Get("/api/results/:id/forms", func(c *fiber.Ctx) error {
		var form Form
		db.Find(&form, "result_id = ?", c.Params("id"))
		return c.JSON(form)
	})


	app.Listen(":8002")
}