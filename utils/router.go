package utils

import (
	"html/template"

	"github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/gin"
)

// InitAssetsTemplates initializes the router to use either a ricebox or the
// filesystem in case the ricebox couldn't be found.
func InitAssetsTemplates(r *gin.Engine, tbox, abox *rice.Box, names ...string) error {
	var err error

	if tbox != nil {
		mt := multitemplate.New()
		var tmpl string
		var message *template.Template
		for _, x := range names {
			if tmpl, err = tbox.String(x); err != nil {
				return err
			}
			if message, err = template.New(x).Parse(tmpl); err != nil {
				return err
			}
			mt.Add(x, message)
		}
		r.HTMLRender = mt
	} else {
		r.LoadHTMLGlob("templates/*")
	}

	if abox != nil {
		r.StaticFS("/static", abox.HTTPBox())
	} else {
		r.Static("/static", "assets")
	}
	return nil
}
