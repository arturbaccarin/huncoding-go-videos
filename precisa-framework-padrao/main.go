package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		json.Unmarshal(r.Body, &w)
		w.Write([]byte("Hello World"))
	})

	mux.HandleFunc("/test/{name}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello " + r.URL.Query().Get("name")))
	})

	http.ListenAndServe(":8080", mux)
}

func testGin() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.ShouldBindBodyWithJSON()
		c.String(200, "Hello World")
	})

	r.GET("/:name", func(c *gin.Context) {
		name := c.Query("name")
		c.String(200, "Hello %s"+name)
	})

	r.Run(":8080")
}
