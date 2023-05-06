package initializers

import "goauth/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
