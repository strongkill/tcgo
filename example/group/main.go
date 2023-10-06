package main

import (
	"github.com/strongkill/tcgo"
	"net/http"
)
func main() {
	r := tcgo.New()
	r.GET("/index", func(c *tcgo.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *tcgo.Context) {
			c.HTML(http.StatusOK, "<h1>Hello tcgo</h1>")
		})

		v1.GET("/hello", func(c *tcgo.Context) {
			// expect /hello?name=tcgo
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *tcgo.Context) {
			// expect /hello/tcgo
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *tcgo.Context) {
			c.JSON(http.StatusOK, tcgo.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":3000")
}
