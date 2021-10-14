package rest

import (
	"kuwait-test/storage"
	"kuwait-test/svc/configs"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type Server struct {
	HostPort string
	storage  storage.Storage
	Log      *logrus.Logger
	c        *configs.Config
}

func NewServer(logger *logrus.Logger, db storage.Storage, c *configs.Config) (*Server, error) {

	return &Server{
		Log:     logger,
		storage: db,
		c:       c,
	}, nil
}

func (s *Server) Run() error {
	r := gin.Default()

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   strings.Split(s.c.CORSHosts, ","),
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))

	pprof.Register(r)

	openAccessed := r.Group("/")
	{
		openAccessed.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"ping": "pong"})
		})
		openAccessed.POST("/accounts", s.CreateAccount)
		openAccessed.GET("/customers/:id/orders", s.ListOrdersByCustomerID)
		openAccessed.GET("/orders/:order_id/customers/:customer_id", s.GetOrderByCustomerID)
		openAccessed.POST("/customers/:customer_id/orders", s.CreateOrder)
		openAccessed.POST("/customers/:customer_id/orders/:order_id/payments", s.PayOrder)
	}

	err := r.Run(s.c.HostPort)
	if err != nil {
		return errors.Errorf("serving on %s failed: %v", s.c.HostPort, err)
	}

	return nil
}
