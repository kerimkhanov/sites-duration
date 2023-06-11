package main

import (
	"github.com/kerimkhanov/sites-duration/internal/app"
	"github.com/kerimkhanov/sites-duration/internal/config"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	lg := logrus.New()
	lg.SetFormatter(&logrus.JSONFormatter{})
	lg.SetOutput(os.Stdout)

	cfg, err := config.InitConfig(lg)
	if err != nil {
		return
	}
	app.Run(cfg, lg)

}
