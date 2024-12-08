package server

import (
	"automatic-doodle/pkg/env"
	"automatic-doodle/pkg/errors"
	"automatic-doodle/types"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Logger interface {
	Trace(format string, args ...any)
	Info(format string, args ...any)
	Warning(format string, args ...any)
	Error(format string, args ...any)
	Fatal(format string, args ...any)
}

type ConfigModule interface {
	GetConfig(string) (string, error)
	GetConfigInt(string) (int, error)
}

type Router interface {
	GetRoutes() *[]types.HandlerItem
}

func New(
	cfg ConfigModule,
	log Logger,
	router Router,
) *Server {
	return &Server{
		cfg:    cfg,
		log:    log,
		router: router,
		app: fiber.New(fiber.Config{
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				if err == nil {
					return nil
				}

				if customError, ok := err.(*errors.CustomError); ok {
					log.Error(
						"Something went wrong during a request: %w",
						customError,
					)
					return ctx.
						Status(customError.Status).
						JSON(customError.Error())
				}

				return ctx.
					Status(fiber.StatusInternalServerError).
					JSON(fiber.Map{
						"message": "Internal Server Error",
					})
			},
		}),
	}
}

type Server struct {
	cfg    ConfigModule
	log    Logger
	router Router
	port   int
	app    *fiber.App
}

func (server *Server) enableCors() {
	server.app.Use(cors.New(
		cors.Config{
			Next:             nil,
			AllowOriginsFunc: nil,
			AllowOrigins:     "*",
			AllowMethods: strings.Join(
				[]string{
					fiber.MethodGet,
					fiber.MethodPost,
					fiber.MethodHead,
					fiber.MethodPut,
					fiber.MethodDelete,
					fiber.MethodPatch,
					fiber.MethodOptions,
				},
				",",
			),
			AllowHeaders:     "",
			AllowCredentials: false,
			ExposeHeaders:    "",
			MaxAge:           0,
		},
	))
}

func (server *Server) injectHandlers(handlers *[]types.HandlerItem) {
	for _, handler := range *handlers {
		switch handler.Method {
		case "GET":
			server.app.Get(handler.Path, handler.Handler...)
		case "POST":
			server.app.Post(handler.Path, handler.Handler...)
		case "PUT":
			server.app.Put(handler.Path, handler.Handler...)
		case "PATCH":
			server.app.Patch(handler.Path, handler.Handler...)
		case "DELETE":
			server.app.Delete(handler.Path, handler.Handler...)
		default:
			server.log.Fatal("Unkown http method %s !", handler.Method)
		}

		server.log.Info("Endpoint injected: %s - %s", handler.Method, handler.Path)
	}
}

func (server *Server) Listen(port int) {
	if env.GO_ENV != types.GoEnvProduction {
		server.enableCors()
	}

	server.injectHandlers(server.router.GetRoutes())

	if err := server.app.Listen(fmt.Sprintf(":%d", port)); err != nil {
		server.log.Fatal("Something went wrong during listen: %v", err)
	}
}
