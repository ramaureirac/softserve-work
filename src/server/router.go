package server

import (
	"net/http"
	"os"

	gin "github.com/gin-gonic/gin"
	logic "github.com/ramaureirac/softserve-work/src/internal/logic"
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

	lst := new(logic.BlockList)

	// populate database
	lst.Add("secure.site", "/notavirus.exe")
	lst.Add("hecker.info", "/dolphin.exe")
	lst.Add("scemer.xyz", "/info/haxx/download")

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"code": "OK"})
		c.Status(http.StatusOK)
	})

	router.GET("/urlinfo/:hostname_and_port/*original_path_and_query_string", func(c *gin.Context) {
		host := c.Param("hostname_and_port")
		query := c.Param("original_path_and_query_string") //+ "?" + c.Request.URL.RawQuery
		status, _ := lst.Search(host, query)
		c.JSON(http.StatusOK, gin.H{"host": host, "query": query, "scan": status})
	})

	router.POST("/urlinfo/:hostname_and_port/*original_path_and_query_string", func(c *gin.Context) {
		host := c.Param("hostname_and_port")
		query := c.Param("original_path_and_query_string")
		err := lst.Add(host, query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{"status": "error"})
			return
		}
		c.JSON(http.StatusCreated, map[string]string{"status": "ok"})
	})

	return router
}
