package gamesapi

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type GamesAPIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *GamesAPIServer {
	return &GamesAPIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func Start(s *GamesAPIServer) error {
	if err:= s.configureLogger(); err!= nil {
		return err
	}

	s.logger.Info("starting games api server")
	s.configureRouter()

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *GamesAPIServer) configureLogger() error {
	level, err:= logrus.ParseLevel(s.config.LogLevel)

	if (err != nil) {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *GamesAPIServer) configureRouter() {
	s.router.HandleFunc("/healthcheck", s.healthcheckHandler())
}

func (s *GamesAPIServer) healthcheckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if (r.Method != http.MethodGet) {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}

		io.WriteString(w, "ok")
	}
}
