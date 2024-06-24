package main

import (
	"log"

	"github.com/SineChat/notification-ms/cronjobs"
	"github.com/SineChat/notification-ms/external/request"
	"github.com/SineChat/notification-ms/internal/config"
	"github.com/SineChat/notification-ms/internal/models/migrations"
	"github.com/SineChat/notification-ms/pkg/repository/storage/mongodb"
	"github.com/SineChat/notification-ms/pkg/router"
	"github.com/SineChat/notification-ms/utility"

	"github.com/go-playground/validator/v10"
)

func main() {
	logger := utility.NewLogger() //Warning !!!!! Do not recreate this action anywhere on the app

	configuration := config.Setup(logger, "./app")

	mongodb.ConnectToDB(logger, configuration.Databases.ConnectionString)

	validatorRef := validator.New()
	db := mongodb.Connection()

	if configuration.Databases.Migrate {
		migrations.RunAllMigrations(db)
	}

	r := router.Setup(logger, validatorRef, db)
	cronjobs.StartCronJob(request.ExternalRequest{Logger: logger}, db, "send-notifications")

	utility.LogAndPrint(logger, "Server is starting at 127.0.0.1:%s", configuration.Server.Port)
	log.Fatal(r.Run(":" + configuration.Server.Port))
}
