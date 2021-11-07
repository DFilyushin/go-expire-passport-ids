package api

import (
	"context"
	_ "expired-passport-checker/docs"
	controllers2 "expired-passport-checker/internal/api/controllers"
	service2 "expired-passport-checker/internal/service"
	"fmt"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"time"
)

type HttpServer struct {
	service    *service2.PassportIdService
	httpServer *http.Server
	Router     *mux.Router
	Host       string
	Port       int
}

func (s *HttpServer) initControllers() {
	api := controllers2.ApiController{}
	infra := controllers2.InfrastructureController{}

	s.Router.HandleFunc("/checkPassport", api.CheckPassport(s.service)).Methods("GET")
	s.Router.HandleFunc("/health", infra.GetHealthCheck()).Methods("GET")
	s.Router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}

func (s *HttpServer) Initialize(serviceHost string, servicePort int, service *service2.PassportIdService) {
	s.Host = serviceHost
	s.Port = servicePort
	s.Router = mux.NewRouter()
	s.service = service
	s.initControllers()
}

func (s *HttpServer) RunService() error {
	/*Запуск http-сервера*/
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	s.httpServer = &http.Server{
		Addr:           addr,
		Handler:        s.Router,
		MaxHeaderBytes: 1 << 20, // 1*2^20 - 128kBytes ~ 1Mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return http.ListenAndServe(addr, s.Router)
}

func (s *HttpServer) Shutdown(ctx context.Context) error {
	/*Остановка*/
	return s.httpServer.Shutdown(ctx)
}
