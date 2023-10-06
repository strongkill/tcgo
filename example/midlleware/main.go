package main

import (
	"github.com/strongkill/tcgo"
	"net/http"
)
func onlyForV2() tcgo.HandlerFunc {
	return func(c *tcgo.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := tcgo.New()
	r.Use(tcgo.Logger()) // global midlleware
	r.GET("/", func(c *tcgo.Context) {
		c.HTML(http.StatusOK, "<h1>Hello tcgo</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *tcgo.Context) {
			// expect /hello/tcgo
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":3000")
}
