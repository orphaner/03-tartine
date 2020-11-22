package httpd

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/orphaner/03-tartine/pkg/core/logger"
	flag "github.com/spf13/pflag"
	"net/http"
	"strconv"
)

var (
	enableGzip bool
	port       int
	Engine     *gin.Engine
)

func InitFlags() {
	flag.BoolVar(&enableGzip, "gzip", true, "enable gzip http compression")
	flag.IntVar(&port, "port", 8080, "httpd port")
}

func InitGinEngine() {
	gin.SetMode(gin.DebugMode)
	Engine = gin.New()
	Engine.HandleMethodNotAllowed = true
	Engine.Use(gin.Recovery())
	Engine.Use(gin.Logger())
	if enableGzip {
		logger.Log.Info("gzip is enabled")
		Engine.Use(gzip.Gzip(gzip.DefaultCompression))
	}
	Engine.GET("/info", infoHandler)
}

func infoHandler(context *gin.Context) {
	context.String(http.StatusOK, "running")
}

func Run() {
	Engine.Run(":" + strconv.Itoa(port))
}
