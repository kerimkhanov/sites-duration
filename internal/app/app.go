package app

import (
	"context"
	"github.com/kerimkhanov/sites-duration/internal/delivery"
	"github.com/kerimkhanov/sites-duration/internal/server"
	"github.com/kerimkhanov/sites-duration/internal/service"
	"github.com/kerimkhanov/sites-duration/internal/storage"
	utils "github.com/kerimkhanov/sites-duration/pkg"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *viper.Viper, lg *logrus.Logger) {
	redisClient, err := storage.InitRedis(cfg)
	if err != nil {
		lg.Println(err)
		return
	}
	defer func() {
		if err := redisClient.Close(); err != nil {
			lg.Printf("can't close Redis client: %v\n", err)
		} else {
			lg.Println("Redis client closed")
		}
	}()

	siteService := service.NewSiteService(redisClient)
	handler := delivery.NewSiteHandler(siteService)

	server := new(server.Server)
	go func() {
		if err := server.Start(cfg, handler.InitRoutes()); err != nil {
			lg.Println("Error starting server:", err)
			return
		}
	}()

	go utils.CheckAndSave(redisClient, cfg, lg)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.TODO(), cfg.GetDuration("APP.SERVER.SHUTDOWN_TIMEOUT"))
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		lg.Println("Error shutting down server:", err)
		return
	}
}
