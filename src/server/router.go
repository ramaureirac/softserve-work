package server

import (
	"net/http"
	"os"

	gin "github.com/gin-gonic/gin"
)

func NewRouterApp() *gin.Engine {

	switch os.Getenv("GIN_MODE") {
	case "testing":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.SetTrustedProxies([]string{"0.0.0.0"})

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"code": "OK"})
		c.Status(http.StatusOK)
	})

	return router
}
