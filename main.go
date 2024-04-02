package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	ginRun()
}

func ginRun() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})
	r.Run(":20011")
}
