package api

import (
	"github.com/gin-gonic/gin"
	"github.com/markitos/markitos-svc-boilerplate/internal/domain"
)

type Server struct {
	address    string
	repository domain.BoilerplateRepository
	router     *gin.Engine
}

func NewServer(address string, repository domain.BoilerplateRepository) *Server {
	apiREST := &Server{
		address:    address,
		repository: repository,
	}
	apiREST.createRouter()

	return apiREST
}

func (s *Server) Repository() domain.BoilerplateRepository {
	return s.repository
}

func (s *Server) Router() *gin.Engine {
	return s.router
}

func (s *Server) createRouter() {
	s.router = gin.Default()
	s.router.POST("/v1/boilerplates", s.create)
	s.router.PATCH("/v1/boilerplates/:id", s.update)
	s.router.DELETE("/v1/boilerplates/:id", s.delete)
	s.router.GET("/v1/boilerplates/:id", s.one)
	s.router.GET("/v1/boilerplates/all", s.all)
	s.router.GET("/v1/boilerplates", s.search)

}

func (s *Server) Run() error {
	return s.router.Run(s.address)
}

func errorResonses(err error) gin.H {
	return gin.H{"error": err.Error()}
}
