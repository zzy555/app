package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	ginRun()
}

func ginRun() {
	r := gin.Default()
	r.GET("/test", test)
	r.Run(":20011")
}

func test(c *gin.Context) {
	num := runtime.NumGoroutine()
	fmt.Println("goroutine:", num)

	n := 100
	for i := 1; i <= n; i++ {
		if i == n {
			fmt.Println("i:", i)
		}
	}

	c.String(http.StatusOK, "test")
}
