package main

import (
	"os"
	"os/signal"
	"context"
	// "syscall"
	"fmt"
	"net/http"
	"log"
	"time"

	"GhdApi/pkg/settings"
	"GhdApi/routers"
	// "github.com/fvbock/endless"
)

func main() {
	// Method 1: 普通使用方法
	// router := routers.InitRouter()
	// s := &http.Server{
	// 	Addr:           fmt.Sprintf(":%d", settings.HTTPPort),
	// 	Handler:        router,
	// 	ReadTimeout:    settings.ReadTimeout,
	// 	WriteTimeout:   settings.WriteTimeout,
	// 	MaxHeaderBytes: 1 << 20,
	// }

	// s.ListenAndServe() // listen and serve on 0.0.0.0:8080

	// Method 3: Gloang > 1.8  http.Server -> shutdown 方式
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", settings.HTTPPort),
		Handler:        router,
		ReadTimeout:    settings.ReadTimeout,
		WriteTimeout:   settings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting...")

	// TODO: 需要重启服务

	// Method 2: endless 热重启方案
	// endless.DefaultReadTimeOut = settings.ReadTimeout
    // endless.DefaultWriteTimeOut = settings.WriteTimeout
    // endless.DefaultMaxHeaderBytes = 1 << 20
    // endPoint := fmt.Sprintf(":%d", settings.HTTPPort)

    // server := endless.NewServer(endPoint, routers.InitRouter())
    // server.BeforeBegin = func(add string) {
    //     log.Printf("Actual pid is %d", syscall.Getpid())
    // }

    // err := server.ListenAndServe()
    // if err != nil {
    //     log.Printf("Server err: %v", err)
    // }
}
