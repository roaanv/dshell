package main
import (
    "net/http"
    _ "embed"
    "github.com/gin-gonic/gin"
	"runtime"
	"fmt"
)

// album represents data about a record album.
func main() {
    router := gin.Default()
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

func arch(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	})
}

func healthcheck(c *gin.Context) {
	//c.Data(http.StatusOK, "application/text", []byte("OK"))
	c.String(http.StatusOK, "OK")
}

