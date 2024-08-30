package app

import (
	"fmt"
	"net"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"

	"github.com/marianozunino/goashot/internal/handler"
	"github.com/marianozunino/goashot/internal/service"
	storage "github.com/marianozunino/goashot/internal/storage/json"
)

type App interface {
	Run()
}

type app struct {
	srv  *echo.Echo
	port string
	host string
}

type ServerOpt func(*app)

func WithPort(port string) ServerOpt {
	return func(p *app) {
		p.port = port
	}
}

func WithHost(host string) ServerOpt {
	return func(p *app) {
		p.host = host
	}
}

func NewBackofficeApp(opts ...ServerOpt) App {
	a := &app{
		port: ":5000",
		host: "localhost",
	}

	for _, opt := range opts {
		opt(a)
	}

	db := storage.NewDB()
	repo := storage.NewRepository(db)
	svc := service.NewOrderService(repo)

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	// Coworkers are evil, so we need to limit the number of requests per second
	e.Use(
		middleware.RateLimiter(
			middleware.NewRateLimiterMemoryStore(
				rate.Limit(5),
			),
		),
	)

	e.Static("/assets", "assets")

	e.GET("/", func(c echo.Context) error {
		http.Redirect(c.Response(), c.Request(), "/orders", http.StatusFound)
		return nil
	})

	handler.OrdersRessources{OrdersService: svc}.Routes(e)

	a.srv = e
	return a
}

func (a *app) Run() {
	hostPort := net.JoinHostPort(a.host, a.port)
	fmt.Println("Listening on port", hostPort)
	a.srv.Logger.Fatal(a.srv.Start(hostPort))
}
