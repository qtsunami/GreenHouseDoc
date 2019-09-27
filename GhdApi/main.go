package main

import (
	"fmt"
	"net/http"

	"GhdApi/pkg/settings"
	"GhdApi/routers"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", settings.HTTPPort),
		Handler:        router,
		ReadTimeout:    settings.ReadTimeout,
		WriteTimeout:   settings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe() // listen and serve on 0.0.0.0:8080
}
