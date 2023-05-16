package migrations

import (
	"github.com/sebastian009w/go_rest/app/models"
	"github.com/sebastian009w/go_rest/platform/database"
)

func Migrate() {
	database.Conn.AutoMigrate(&models.Post{})
}
