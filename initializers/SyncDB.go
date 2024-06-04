package initializers

import (
	"/models"
	"fmt"
)

func SyncDB() {
	DB.AutoMigrate(&models.StudentModel{})
	fmt.Println("SyncedDB")
}
