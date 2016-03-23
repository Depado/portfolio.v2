package main

import (
	"log"
	"net/http"

	"github.com/Depado/portfolio/utils"
	"github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
)

func main() {
	var err error

	tbox, _ := rice.FindBox("templates")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	if err = utils.InitAssetsTemplates(r, tbox, nil, "index.html"); err != nil {
		log.Fatal(err)
	}
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.Run(":8005")
}
