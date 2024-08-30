package app

// import (
// 	"context"
//
// 	"github.com/marianozunino/goashot/internal/config"
// 	"github.com/marianozunino/goashot/internal/service"
// 	storage "github.com/marianozunino/goashot/internal/storage/json"
// 	"go.uber.org/fx"
// )

// var App = fx.New(
// 	storage.Module,
// 	service.Module,
// 	config.Module,
// 	fx.Invoke(
// 		func(lifecycle fx.Lifecycle, scraper service.ScrapperService) {
// 			lifecycle.Append(
// 				fx.Hook{
// 					OnStart: func(context.Context) error {
// 						go scraper.Scrape()
// 						return nil
// 					},
// 				},
// 			)
// 		},
// 	),
// )
