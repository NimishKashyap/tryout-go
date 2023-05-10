package migrations

import (
	"log"

	"github.com/NimishKashyap/tryout-go/api/db"
	"github.com/NimishKashyap/tryout-go/pkg/models"
	_ "gorm.io/driver/postgres"
)

func Migrate() {
	database := db.GetDB()
	log.Println("PERFORMING MIGRATIONS")
	database.Raw("CREATE EXTENSIION IF NOT EXISTS \"uuid-ossp\";")
	if err := database.AutoMigrate(&models.Users{}, &models.BaseModel{}); err != nil {
		log.Fatal(err)
	}
}
