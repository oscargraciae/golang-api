package api

import (
	"api-app/api/config"
	"api-app/api/router"

	"github.com/gin-gonic/gin"
)

// Run server
func Run() {
	r := gin.Default()
	router.SetupRouter(r)
	if true {
		config.Load()
	}
	r.Run(":9000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
