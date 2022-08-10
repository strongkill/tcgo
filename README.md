# tcgo
web framework writen by Go language

##HOW TO USE

```go
#add in go.mod

require (
	github.com/strongkill/tcgo v1.0.2
)

```

###example 1
```go
package main

import (
	"github.com/strongkill/tcgo/tcgo"
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

```
###example 2
```go
package main

import (
	"github.com/strongkill/tcgo/tcgo"
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
```
###example midlleware
```go
package main

import (
	"github.com/strongkill/tcgo/tcgo"
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

	r.Run(":9999")
}
```

###example templates
```go
package main

import (
	"github.com/strongkill/tcgo/tcgo"
	"net/http"
)
type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := tcgo.New()
	r.Use(tcgo.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "tcgoktutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *tcgo.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *tcgo.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", tcgo.H{
			"title":  "tcgo",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *tcgo.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", tcgo.H{
			"title": "tcgo",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
}
```
