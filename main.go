package main

import (
	"fmt"
	"gin_log/pkg/setting"
	"gin_log/routers"
	"net/http"

	_ "gin_log/models"
)

func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.HTTPPort),
		Handler: router,
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}
