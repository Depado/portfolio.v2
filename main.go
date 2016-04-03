package main

import (
	"log"
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"

	"github.com/Depado/portfolio.v2/fetch"
	"github.com/Depado/portfolio.v2/utils"
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"services":   fetch.Current,
		"monit_root": fetch.MONITROOT,
	})
}

func ascii(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{
		"services":   fetch.Current,
		"monit_root": fetch.MONITROOT,
	})
}

func main() {
	var err error

	go fetch.Start()

	tbox, _ := rice.FindBox("templates")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	if err = utils.InitAssetsTemplates(r, tbox, nil, "index.html", "about.html"); err != nil {
		log.Fatal(err)
	}
	r.GET("/", index)
	r.GET("/about", ascii)
	r.Run(":8005")
}
