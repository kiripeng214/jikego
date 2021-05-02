package xnet

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

func InitServer() error {
	srv := initRouter()
	go func() {
		srv.ListenAndServe()
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-quit:
		fmt.Println("shut down the server")
	}
	//关闭其他服务
	if err := srv.Shutdown(context.TODO()); err != nil {
		fmt.Println("Server Shutdown: ", err)
		return err
	}
	fmt.Println("Server Shutdown: success")
	return nil
}
