package main

import (
	"github.com/strongkill/tcgo"
	"net/http"
)

func main() {
	r := tcgo.New()
	r.GET("/", func(c *tcgo.Context) {
		c.HTML(http.StatusOK, "<h1>Hello tcgo</h1>")
	})

	r.GET("/hello", func(c *tcgo.Context) {
		// expect /hello?name=tcgo
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *tcgo.Context) {
		// expect /hello/tcgo
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *tcgo.Context) {
		c.JSON(http.StatusOK, tcgo.H{"filepath": c.Param("filepath")})
	})

	r.Run(":3000")
}
