package app

import (
	"context"
	"os"

	"github.com/marianozunino/goashot/internal/config"
	"github.com/marianozunino/goashot/internal/http"
	"github.com/marianozunino/goashot/internal/service"
	storage "github.com/marianozunino/goashot/internal/storage/json"
	"go.uber.org/fx"
)

var App = fx.New(
	storage.Module,
	service.Module,
	config.Module,
	http.Module,
	fx.Invoke(registerHooks),
)

func registerHooks(lifecycle fx.Lifecycle, gin *http.GinHandler) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				// read port from env "PORT"
				// if not set, use default port 8080
				var port = os.Getenv("PORT")
				if port == "" {
					port = "5000"
				}
				go gin.Run(":" + port)
				return nil
			},
		},
	)
}
