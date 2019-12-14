package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerRoutes() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
		// c.String(http.StatusOK, "Hello from %v", "Gin")
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.GET("/employees/:id/vacation", func(c *gin.Context){
		id := c.Param("id")
		timesOff, ok := TimesOff[id]

		if !ok{
			c.String(http.StatusNotFound, "404 - Page Not Found")
			return
		}

		c.HTML(http.StatusOK, "vacation-voewview.html",
			map[string]interface{}{
				"TimesOff": timesOff,
			})

	})

	admin := r.Group("/admin")

	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html", nil)
	})

	r.Static("/public", "./public")
	// r.StaticFS("/public", http.Dir("./public"))

	return r
}
