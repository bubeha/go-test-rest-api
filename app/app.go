package app

import (
	"net/http"

	"github.com/bubeha/go-test-rest-api/router"
	"github.com/sirupsen/logrus"
)

// IApplication ..
type IApplication interface {
	Init()
	Start() error
}

// Application ...
type Application struct {
	router router.IRouter
	logger *logrus.Logger
}

// Init ...
func (application *Application) Init() {
	application.router = router.New()
	application.router.InitRoutes()
}

// Start ..
func (application *Application) Start() error {
	application.logger.Infoln("Server (http://127.0.0.1:8000) started")

	return http.ListenAndServe(":8000", application.router.RouteMultiplexer())
}

// New ...
func New() IApplication {
	return &Application{
		logger: logrus.New(),
	}
}
