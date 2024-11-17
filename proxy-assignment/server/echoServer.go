package server

import (
	"fmt"

	"inwpuun/proxy_assignment/config"
	cryptoDomain "inwpuun/proxy_assignment/internal/crypto/domain"
	cryptoHandlers "inwpuun/proxy_assignment/internal/crypto/handlers"
	cryptoRepositories "inwpuun/proxy_assignment/internal/crypto/repositories"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type echoServer struct {
	app  *echo.Echo
	conf *config.Config
}

func NewEchoServer(conf *config.Config) Server {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	return &echoServer{
		app:  echoApp,
		conf: conf,
	}
}

func (s *echoServer) Start() {
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())

	s.app.GET("v1/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	s.initializeCryptoHttpHandler()

	serverUrl := fmt.Sprintf(":%d", s.conf.ServerConfig.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func (s *echoServer) initializeCryptoHttpHandler() {
	// Initialize all layers
	cryptoRepository := cryptoRepositories.NewCryptoRepository(s.conf)

	cryptoUsecase := cryptoDomain.NewCryptoDomainImpl(cryptoRepository)

	cryptoHttpHandler := cryptoHandlers.NewCryptoHttpHandler(cryptoUsecase)

	// Routers
	cryptoRouters := s.app.Group("v1/crypto")
	cryptoRouters.POST("/broadcast", cryptoHttpHandler.BroadcastTransaction)
	cryptoRouters.GET("/check/:hash", cryptoHttpHandler.CheckStatus)
}
