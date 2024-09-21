package main

import (
	"clickercopro/handlers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var point = handlers.Point{}

func main() {

	router := gin.Default()

	router.Static("/assets", "./assets")
	router.StaticFile("/assets", "./assets")

	router.GET("/", func(c *gin.Context) {
		if handlers.ValidAuth(c.ClientIP()) == true {
			c.Redirect(http.StatusFound, "/index")
		} else {
			c.Redirect(http.StatusFound, "/login")

		}

	})
	router.GET("/index", func(ctx *gin.Context) {
		ctx.File("html/index.html")
	})

	router.GET("/login", func(c *gin.Context) {
		c.File("html/login.html")
	})

	router.POST("/", func(ctx *gin.Context) {
		username, password := ctx.PostForm("username"), ctx.PostForm("password")

		fmt.Println(username, password)

		if username != "" && password != "" {
			ctx.Redirect(http.StatusFound, "/index")
		}
	})

	router.POST("/endpoint", func(ctx *gin.Context) {
		point.AddPoint()
		fmt.Println(point)
		ctx.JSON(http.StatusOK, point)
	})

	router.Run(":3000")
}
