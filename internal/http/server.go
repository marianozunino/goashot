package http

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"go.uber.org/fx"
)

type GinHandler struct {
	*gin.Engine
}

var Module = fx.Options(
	fx.Provide(registerGinHandler),
	fx.Invoke(registerRoutes),
)

func registerGinHandler() *GinHandler {
	// Coworkers are evil, so we need to limit the number of requests per second
	rate := limiter.Rate{
		Period: 1 * time.Second,
		Limit:  5,
	}
	store := memory.NewStore()
	limiterInstance := limiter.New(store, rate)
	middleware := mgin.NewMiddleware(limiterInstance)

	gin.SetMode(gin.ReleaseMode)

	instance := gin.Default()

	instance.LoadHTMLGlob("templates/*.tmpl")
	instance.Use(middleware)

	handler := GinHandler{
		instance,
	}

	return &handler
}
