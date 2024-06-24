package router

import (
	"net/http"

	"github.com/SineChat/notification-ms/pkg/middleware"
	"github.com/SineChat/notification-ms/pkg/repository/storage/mongodb"
	"github.com/SineChat/notification-ms/utility"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Setup(logger *utility.Logger, validator *validator.Validate, db *mongodb.Database) *gin.Engine {
	r := gin.New()

	// Middlewares
	// r.Use(gin.Logger())
	r.Use(middleware.Logger())
	r.Use(middleware.Security())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.MaxMultipartMemory = 1 << 20 // 1MB

	ApiVersion := "v1"
	Health(r, ApiVersion, validator, db, logger)
	Notifications(r, ApiVersion, validator, db, logger)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"name":    "Not Found",
			"message": "Page not found.",
			"code":    404,
			"status":  http.StatusNotFound,
		})
	})

	return r
}
