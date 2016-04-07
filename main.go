package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Depado/portfolio.v2/fetch"
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"services":   fetch.Current,
		"monit_root": fetch.MONITROOT,
	})
}

func about(c *gin.Context) {
	c.HTML(http.StatusOK, "about.tmpl", gin.H{})
}

func main() {

	go fetch.Start()

	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Static("/static", "assets")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", index)
	r.GET("/about", about)

	r.Run(":8005")
}
