package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	api "wangzheng/brain/cmd/server/api"

	"github.com/gin-gonic/gin"
)

const DefaultPort = ":8080"

func main() {
	router := gin.Default()
	api.InitApis(router)

	srv := &http.Server{
		Addr: ":8000",
		// ReadHeaderTimeout is the amount of time allowed to read
		// request headers. The connection's read deadline is reset
		// after reading the headers and the Handler can decide what
		// is considered too slow for the body. If ReadHeaderTimeout
		// is zero, the value of ReadTimeout is used. If both are
		// zero, there is no timeout.
		ReadHeaderTimeout: 15 * time.Second,

		// ReadTimeout is the maximum duration for reading the entire
		// request, including the body. A zero or negative value means
		// there will be no timeout.
		//
		// Because ReadTimeout does not let Handlers make per-request
		// decisions on each request body's acceptable deadline or
		// upload rate, most users will prefer to use
		// ReadHeaderTimeout. It is valid to use them both.
		ReadTimeout: 15 * time.Second,

		// WriteTimeout is the maximum duration before timing out
		// writes of the response. It is reset whenever a new
		// request's header is read. Like ReadTimeout, it does not
		// let Handlers make decisions on a per-request basis.
		// A zero or negative value means there will be no timeout.
		WriteTimeout: 10 * time.Second,

		// IdleTimeout is the maximum amount of time to wait for the
		// next request when keep-alives are enabled. If IdleTimeout
		// is zero, the value of ReadTimeout is used. If both are
		// zero, there is no timeout.
		IdleTimeout: 30 * time.Second,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
