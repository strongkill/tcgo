package main

import (
	"net/http"
	"tcgo"
)

func main() {
	r := tcgo.Default()
	r.GET("/", func(c *tcgo.Context) {
		c.String(http.StatusOK, "Hello world\n")
	})
	r.Run(":3000")
}
