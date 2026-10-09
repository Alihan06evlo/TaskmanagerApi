package apiserver

import (
	"net/http"
	task "taskmanagerApi/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	router *gin.Engine
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: gin.Default(),
	}
}
func (s *APIServer) Run() error {
	logrus.WithField("addr", s.config.BindAddr).Info("[Server] Starting HTTP server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}
func (s *APIServer) ConfigureRouter(h *task.Handler) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.HandleMethodNotAllowed = true
	r.POST("/tasks", h.Create)
	r.DELETE("/tasks/:id", h.Delete)

	s.router = r
}
