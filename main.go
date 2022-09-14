package main

import (
	"context"
	"fmt"
	"gin_log/models"
	_ "gin_log/models"
	"gin_log/pkg/logging"
	"gin_log/pkg/setting"
	"gin_log/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
}
func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Println("listen:", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Server Shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Println("Shutdown:", err)
	}

	log.Println("Server Exit")
}
