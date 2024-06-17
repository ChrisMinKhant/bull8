package car

import (
	"log"

	"github.com/ChrisMinKhant/megoyougo_framework/util"
	"gorm.io/gorm"
)

/*
 * Example ' Car ' entity
 */

var dbHelper = util.GetDbHelperInstance()

type Car struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

func New() *Car {
	return &Car{}
}

/*
 * This function is important because it will
 * be called from database to migrate this
 * entity.
 */

func (car *Car) Initialize(dbConnection *gorm.DB) {
	dbConnection.AutoMigrate(&Car{})
}

func (car *Car) Create() {

	db := dbHelper.Connect()

	result := db.Create(&Car{
		Brand: "Toyota",
		Model: "HighAce",
		Year:  2024,
	})

	log.Printf("Affected row ::: %v\n", result.RowsAffected)
}
