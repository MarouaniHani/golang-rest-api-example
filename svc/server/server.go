package server

import (
	"kuwait-test/internal"
	"kuwait-test/storage"
	"kuwait-test/storage/sql"
	"kuwait-test/svc/configs"
	"kuwait-test/svc/rest"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Server struct {
	logger     logrus.FieldLogger
	db         storage.Storage
	restServer *rest.Server
}

func NewServer(c *configs.Config) (*Server, error) {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	db, err := sql.Open(c.Dsn, filepath.Join(internal.GetProjectPath(), "storage/migration"))
	if err != nil {
		return nil, err
	}

	restServer, err := rest.NewServer(logger,db, c)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create a new rest server")
	}

	return &Server{
		logger:     logger,
		restServer: restServer,
	}, nil
}

func (s *Server) Run() error {

	if s.restServer != nil {
		err := s.restServer.Run()
		if err != nil {
			return err
		}
	} else {
		s.logger.Infoln("no rest server started")
	}

	return nil
}
