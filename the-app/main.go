package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)

// album represents data about a record album.
func main() {
	router := gin.Default()
	router.GET("/", home(router))
	router.GET("/ping", ping)
	router.GET("/arch", arch)
	router.GET("/healthcheck", healthcheck)

	router.Run("0.0.0.0:8080")
}

func ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func home(router *gin.Engine) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var routeInfo []string
		for _, aRoute := range router.Routes() {
			routeInfo = append(routeInfo, fmt.Sprintf("%s %s", aRoute.Method, aRoute.Path)) // note the = instead of :=
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"routes": routeInfo,
		})
	}

	return gin.HandlerFunc(fn)
}

func arch(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	})
}

func healthcheck(c *gin.Context) {
	//c.Data(http.StatusOK, "application/text", []byte("OK"))
	c.String(http.StatusOK, "OK")
}
