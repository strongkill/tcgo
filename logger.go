package tcgo

/**
Logger
tcgo a web framework writen by Go language
Author Wing K.Y
2023-02-25
*/
import (
	"github.com/strongkill/goConsole/console"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		console.Log("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
