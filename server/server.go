package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/kimzey/iskeai-shop/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type echoServer struct{
	app *echo.Echo
	db *gorm.DB
	conf *config.Config
}

var(
	once sync.Once
	server *echoServer
)

func NewEchoServer(conf *config.Config , db *gorm.DB) *echoServer{
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	once.Do(func() {
		server = &echoServer{
			app: echoApp,
			db : db,
			conf : conf,
		}
	})

	return server
}

func (s *echoServer) Start(){
	timeOutMideleware := getTimeOutMiddlwware(s.conf.Server.Timeout)
	corsMiddleware := getCORSMiddleware(s.conf.Server.AllowOrigins)
	bodyLimitMiddleware := getBodyLimitMiddleware(s.conf.Server.BodyLimit)
	
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())
	s.app.Use(corsMiddleware)
	s.app.Use(bodyLimitMiddleware)
	s.app.Use(timeOutMideleware)


	s.app.GET("/v1/health", s.healthCheck)
	s.initItemShopRouter()
	
	quitCh := make(chan os.Signal,1)
	
	signal.Notify(quitCh,syscall.SIGINT,syscall.SIGTERM)

	go s.gracefullyShutdown(quitCh)

	s.httpListening()
}

func (s *echoServer) httpListening (){
	Url := fmt.Sprintf(":%d",s.conf.Server.Port)

	err := s.app.Start(Url)

	if err != nil && err != http.ErrServerClosed{
		s.app.Logger.Fatalf("Error: %s",err.Error())
	}
}

func (s *echoServer) gracefullyShutdown (quitCh chan os.Signal) {
	ctx := context.Background()
	// s.app.Logger.Info("test : ",ctx)

	<-quitCh
	s.app.Logger.Info("Shuttdown server")

	if err := s.app.Shutdown(ctx); err != nil {
		s.app.Logger.Fatalf("error %s",err.Error())
	}

}

func (s *echoServer) healthCheck (c echo.Context) error {
	return c.String(http.StatusOK,"Ok")
}

func getLoggerMiddleware () echo.MiddlewareFunc {
	return middleware.Logger()
}

func getTimeOutMiddlwware(timeout time.Duration) echo.MiddlewareFunc {
	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper: middleware.DefaultSkipper,
		ErrorMessage: "request Timeout",
		Timeout: timeout * time.Second,
	})
}

func getCORSMiddleware(allowOrigins []string) echo.MiddlewareFunc{
	return middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper: middleware.DefaultSkipper,
		AllowOrigins: allowOrigins,
		AllowMethods: []string{echo.GET,echo.POST,echo.PUT,echo.PATCH,echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin,echo.HeaderContentType,echo.HeaderAccept},
	})
}

func getBodyLimitMiddleware(bodyLuimit string) echo.MiddlewareFunc{
	return middleware.BodyLimit(bodyLuimit)
}