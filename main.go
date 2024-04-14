package main

import (
	"log"
	"net/http"
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.LoadHTMLFiles("static/index.html")

    e.GET("*path", func(c *gin.Context) {
    	c.HTML(http.StatusOK, "index.html", gin.H{})
		fmt.Println("HTTPS request received!")
 	})

	return e
}

func main() {
	server01 := &http.Server{
		Addr:         ":443",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServeTLS("cert.pem", "key.pem")
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}