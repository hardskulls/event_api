package server

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type AppHandler func(*Server, *fiber.Ctx) error
type Handler func(*fiber.Ctx) error

func applyHandler(app *Server, f AppHandler) Handler {
	res := func(ctx *fiber.Ctx) error {
		return f(app, ctx)
	}
	return res
}

type Server struct {
	port   int
	server *fiber.App
	db     *sql.DB
}

func New(port int) *Server {
	app := fiber.New()
	db, err := sql.Open("clickhouse", "tcp://localhost:9000?debug=true")
	if err != nil {
		panic(err)
	}
	return &Server{port: port, server: app, db: db}
}

func (s *Server) Handle(method, path string, f AppHandler) {
	_ = s.server.Add(method, path, applyHandler(s, f))
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
