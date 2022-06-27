package rest

import (
	"bkind_pet_shelter/src/backend/petCadaster_service/domain"

	"github.com/gin-gonic/gin"
)

type ServerParams struct {
	Port uint32
	App  domain.Application
}

type Server struct {
	port uint32
	app  domain.Application
}

func NewServer(params ServerParams) (Server, error) {
	if params.Port == 0 {
		return Server{}, domain.ErrMissingPortDependency
	}

	if params.App == nil {
		return Server{}, domain.ErrMissingAppDependency
	}

	return Server{port: params.Port, app: params.App}, nil
}

func MustNewServer(params ServerParams) Server {
	s, err := NewServer(params)
	if err != nil {
		panic(err)
	}
	return s
}

func (s *Server) Run() error {
	r := gin.Default()
	r.GET("/ping", s.PingEndpoint(*gin.Context))
	r.Run()
	return nil
}
