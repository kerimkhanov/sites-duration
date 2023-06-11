package server

import (
	"context"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

type Server struct {
	srv http.Server
}

func (s *Server) Start(config *viper.Viper, handlers http.Handler) error {
	s.srv = http.Server{
		Addr:         ":" + config.GetString("APP.SERVER.PORT"),
		Handler:      handlers,
		WriteTimeout: time.Second * time.Duration(config.GetInt("APP.SERVER.WRITE_TIMEOUT")),
		ReadTimeout:  time.Second * time.Duration(config.GetInt("APP.SERVER.READ_TIMEOUT")),
	}
	log.Printf("Server starting http://localhost:%s\n", config.GetString("APP.SERVER.PORT"))
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
