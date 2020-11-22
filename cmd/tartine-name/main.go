package main

import (
	"github.com/gin-gonic/gin"
	"github.com/orphaner/03-tartine/pkg/core/httpd"
	"github.com/orphaner/03-tartine/pkg/core/logger"
	flag "github.com/spf13/pflag"
	"net/http"
)

var (
	name string
)

type NameResult struct {
	Name string `json:"name"`
}

func init() {
	logger.InitFlags()
	httpd.InitFlags()

	flag.StringVar(&name, "name", "raymond", "your name")
	flag.Parse()

	logger.InitLogger()
	httpd.InitGinEngine()
}

func main() {
	httpd.Engine.GET("/name", getNameHandler)
	httpd.Run()
}

func getNameHandler(c *gin.Context) {
	c.JSON(http.StatusOK, NameResult{Name: name})
}
