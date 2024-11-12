package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Service struct {
	DB         *gorm.DB
	HttpRouter *gin.Engine
}

func New() (*Service, error) {
	server := &Service{}

	return server, nil

}

func (s *Service) InitHttpRouter() *Service {
	if s.HttpRouter == nil {
		router := gin.New()
		router.Use(gin.Logger())
		router.Use(gin.Recovery())
		s.HttpRouter = router
		s.setBaseHttpRoutes()
	}
	return s
}
func (s *Service) setBaseHttpRoutes() {
	s.HttpRouter.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})

	s.HttpRouter.GET("/health", func(ctx *gin.Context) {
		// Send a ping to make sure the database connection is alive.

		ctx.JSON(http.StatusOK, gin.H{"live": "ok"})
	})
}
func (s *Service) RunServer() {
	
	if s.HttpRouter == nil {
		s.InitHttpRouter()
	}

	
	err := s.HttpRouter.Run(":8080")
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
