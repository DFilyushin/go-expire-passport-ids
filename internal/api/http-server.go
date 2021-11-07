package api

import (
	"context"
	_ "expired-passport-checker/docs"
	"expired-passport-checker/internal/api/controllers"
	"expired-passport-checker/internal/service"
	"fmt"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"time"
)

const (
	MaxHeaderBytes = 1 << 20 // 1*2^20 - 128kBytes ~ 1Mb
	ReadTimeout    = 10 * time.Second
	WriteTimeout   = 10 * time.Second
)

type HttpServer struct {
	service    *service.PassportIdService
	httpServer *http.Server
	Router     *mux.Router
	Host       string
	Port       int
}

// initControllers инициализирует контроллеры
func (s *HttpServer) initControllers() {
	api := controllers.ApiController{}
	infra := controllers.InfrastructureController{}

	s.Router.HandleFunc("/checkPassport", api.CheckPassport(s.service)).Methods("GET")
	s.Router.HandleFunc("/health", infra.GetHealthCheck()).Methods("GET")
	s.Router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}

func (s *HttpServer) Initialize(serviceHost string, servicePort int, service *service.PassportIdService) {
	s.Host = serviceHost
	s.Port = servicePort
	s.Router = mux.NewRouter()
	s.service = service
	s.initControllers()
}

// RunServer запускает http-сервер
func (s *HttpServer) RunServer() error {
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	s.httpServer = &http.Server{
		Addr:           addr,
		Handler:        s.Router,
		MaxHeaderBytes: MaxHeaderBytes,
		ReadTimeout:    ReadTimeout,
		WriteTimeout:   WriteTimeout,
	}
	return http.ListenAndServe(addr, s.Router)
}

// Shutdown останавливает сервис
func (s *HttpServer) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
