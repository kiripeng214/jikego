package main

import (
	"github.com/gin-gonic/gin"
	"kiripeng214/jikego/week3/xnet"
	"log"
	"net/http"
)

func initRouter() *http.Server {

	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})
	srv := &http.Server{
		Handler: router,
	}

	return srv
}

func main() {
	kiApp := xnet.New(xnet.HttpSever(initRouter()))
	if err := kiApp.Run(); err != nil {
		log.Println(err)
	}
}
