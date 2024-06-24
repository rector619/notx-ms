package migrations

import "github.com/SineChat/notification-ms/internal/models"

// _ = db.AutoMigrate(MigrationModels()...)
func AuthMigrationModels() []interface{} {
	return []interface{}{
		&models.NotificationRecord{},
	}
}
