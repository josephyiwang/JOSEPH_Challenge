package main

import (
	"log"
	"net/http"
	"time"
	"fmt" 
	"strings"

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

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("*path", func(c *gin.Context) {
		reqhost := strings.Split(c.Request.Host, ":")[0]
        c.Redirect(302, "https://" + reqhost + ":443" + c.Request.RequestURI)
		fmt.Println("HTTP request received! Redirecting to https://" + reqhost + ":443" + c.Request.RequestURI)
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

	server02 := &http.Server{
		Addr:         ":8080",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServeTLS("certfile.pem", "keyfile.pem")
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}