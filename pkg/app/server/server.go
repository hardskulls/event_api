package server

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Port = int
type Method = string
type Path = string
type Handler func(*fiber.Ctx) error

type Config struct {
	DB *sql.DB
}

type Route struct {
	method  Method
	path    Path
	handler Handler
}

func NewRoute(method, path string, handler func(c *fiber.Ctx) error) Route {
	return Route{method: method, path: path, handler: handler}
}

type Server struct {
	port   Port
	server *fiber.App
}

func New(port Port) *Server {
	app := fiber.New()
	return &Server{port: port, server: app}
}

func (s *Server) Handle(m Method, p Path, h Handler) {
	_ = s.server.Add(m, p, h)
}

func (s *Server) Add(arg interface{}) {
	_ = s.server.Use(arg)
}

func (s *Server) MustRun() {
	err := s.server.Listen(fmt.Sprintf(":%d", s.port))
	if err != nil {
		panic(err)
	}
}
