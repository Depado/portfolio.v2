package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Depado/portfolio.v2/conf"
	"github.com/Depado/portfolio.v2/fetch"
)

func main() {
	var err error

	if err = conf.Load("conf.yml"); err != nil {
		log.Fatal(err)
	}

	go fetch.Start()

	if !conf.C.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Static("/static", "assets")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"services":   fetch.Current,
			"monit_root": fetch.MONITROOT,
		})
	})
	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.tmpl", gin.H{})
	})

	r.Run(fmt.Sprintf(":%d", conf.C.Port))
}
