package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"runtime"
)

func init() {
    lvl, ok := os.LookupEnv("LOG_LEVEL")
    // LOG_LEVEL not set, let's default to debug
    if !ok {
        lvl = "debug"
    }
    // parse string, this is built-in feature of logrus
    ll, err := log.ParseLevel(lvl)
    if err != nil {
        ll = log.DebugLevel
    }
    // set global log level
    log.SetLevel(ll)
}


// album represents data about a record album.
func main() {
	router := gin.Default()
	router.GET("/", home(router))
	router.GET("/ping", ping)
	router.GET("/arch", arch)
	router.GET("/healthcheck", healthcheck)
	router.GET("/probe/live", liveProbe)
	router.GET("/probe/ready", readyProbe)
	router.GET("/probe/startup", startupProbe)

	router.Run("0.0.0.0:8080")
}

func ping(c *gin.Context) {
	log.Debugln(fmt.Sprintf("called '/%s'", c.Request.URL.Path))
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func home(router *gin.Engine) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		log.Debugln(fmt.Sprintf("called '/%s'", c.Request.URL.Path))
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
	log.Debugln(fmt.Sprintf("called '/%s'", c.Request.URL.Path))
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	})
}

func healthcheck(c *gin.Context) {
	log.Debugln(fmt.Sprintf("called '/%s'", c.Request.URL.Path))
	c.String(http.StatusOK, "OK")
}

func liveProbe(c *gin.Context) {
	log.Debugln(fmt.Sprintf("called '/%s'", c.Request.URL.Path))
	c.String(http.StatusOK, "OK")
}

func readyProbe(c *gin.Context) {
	log.Debugln(fmt.Sprintf("called '/%s'", c.Request.URL.Path))
	c.String(http.StatusOK, "OK")
}

func startupProbe(c *gin.Context) {
	log.Debugln(fmt.Sprintf("called '/%s'", c.Request.URL.Path))
	c.String(http.StatusOK, "OK")
}
