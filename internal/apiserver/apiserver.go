package apiserver

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *http.ServeMux
}

func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: http.NewServeMux(),
	}
}

func (s *APIserver) Start() error {
	if err := s.ConfigureLogger(); err != nil {
		return err
	}

	s.ConfigureRouter()
	s.logger.Info("Starting API server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIserver) ConfigureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIserver) ConfigureRouter() {
	s.router.HandleFunc("/hello", s.HandleHello())
	s.router.HandleFunc("/", s.HandleMain())
	s.router.HandleFunc("/md2html", s.HandleMDFile())
}
