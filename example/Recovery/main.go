package main

import (
	"net/http"
	"github.com/strongkill/tcgo"
)

func main() {
	r := tcgo.Default()
	r.GET("/", func(c *tcgo.Context) {
		c.String(http.StatusOK, "Hello tcgo\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *tcgo.Context) {
		names := []string{"tcgo"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":3000")
}
