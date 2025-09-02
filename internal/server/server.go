package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/aaanu/backendusers/internal/config"
	"github.com/aaanu/backendusers/internal/domain/models"
	"github.com/aaanu/backendusers/internal/domain/requests"
	"github.com/gin-gonic/gin"
)

type UsersService interface {
	User(ctx context.Context, tgID string) (*models.User, error)
	Users(ctx context.Context, role string) ([]models.User, error)
	SaveUser(ctx context.Context, userRequest *requests.SaveUserRequest) error
	UpdateUser(ctx context.Context, tgID string, role string, language string) error
	DeleteUser(ctx context.Context, tgID string, fromUserID string) error
}

type Server struct {
	server *http.Server
	engine *gin.Engine
}

func New(
	service UsersService,
) *Server {
	cfg := config.Config().Server
	engine := gin.Default()

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: engine,
	}

	group := engine.Group("/api")
	group.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	Register(group, service)

	return &Server{
		server: httpServer,
		engine: engine,
	}
}

func (s *Server) Start() {
	if err := s.server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func (s *Server) GracefulStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		panic(err)
	}
}
