package http

import (
	status "net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marianozunino/goashot/internal/config"
	"github.com/marianozunino/goashot/internal/dto"
	"github.com/marianozunino/goashot/internal/service"
)

type RouteHandler struct {
	*gin.Engine
	service service.OrderService
	config  config.Config
}

func registerRoutes(router *GinHandler, service service.OrderService, config config.Config) {
	routeHandler := RouteHandler{
		router.Engine,
		service,
		config,
	}

	routeHandler.Static("/assets", "./public")

	routeHandler.NoRoute(func(c *gin.Context) {
		c.HTML(status.StatusNotFound, "error.tmpl", gin.H{
			"Message": "Where you think you're going?",
		})
	})

	routeHandler.ordersRoutes()
}

func (r *RouteHandler) ordersRoutes() {
	r.GET("/", func(c *gin.Context) {
		c.Redirect(status.StatusFound, "/orders")
	})

	r.GET("/orders", func(c *gin.Context) {
		c.HTML(status.StatusOK, "index.tmpl", gin.H{
			"Orders": r.service.GetOrders(),
		})
	})

	r.GET("/orders/:id", func(c *gin.Context) {
		stringId := c.Param("id")
		id, err := strconv.Atoi(stringId)

		if err != nil {
			c.HTML(status.StatusBadRequest, "error.tmpl", gin.H{
				"Message": ":( ID is not a number",
			})
		}

		if order, err := r.service.GetOrder(id); err != nil {
			c.HTML(status.StatusNotFound, "error.tmpl", gin.H{
				"Message": ":( Order not found",
			})
		} else {
			c.HTML(status.StatusOK, "view.tmpl", gin.H{
				"Order":    order,
				"Proteins": dto.Proteins.GetProteins(),
				"Toppings": dto.Toppings.GetToppings(),
			})
		}
	})

	r.GET("/orders/:id/edit", func(c *gin.Context) {
		stringId := c.Param("id")
		id, err := strconv.Atoi(stringId)

		if err != nil {
			c.HTML(status.StatusBadRequest, "error.tmpl", gin.H{
				"Message": "FUCK! Stop it!",
			})
		}

		if order, err := r.service.GetOrder(id); err != nil {
			c.HTML(status.StatusNotFound, "error.tmpl", gin.H{
				"Message": "Woopsie, order not found",
			})
		} else {
			c.HTML(status.StatusOK, "edit.tmpl", gin.H{
				"Order":    order,
				"Proteins": dto.Proteins.GetProteins(),
				"Toppings": dto.Toppings.GetToppings(),
			})
		}
	})

	r.POST("/orders/:id/edit", func(c *gin.Context) {
		data := &dto.Order{}

		if err := c.ShouldBind(data); err != nil {
			c.HTML(status.StatusBadRequest, "error.tmpl", gin.H{
				"Message": "Why you do this to me?",
			})
		}

		stringId := c.Param("id")
		id, err := strconv.Atoi(stringId)

		if err != nil {
			c.HTML(status.StatusBadRequest, "error.tmpl", gin.H{
				"Message": "Stop trying to hack me",
			})
		}

		if data.ID != id {
			c.HTML(status.StatusBadRequest, "error.tmpl", gin.H{
				"Message": "Really? come on..",
			})
		}

		r.service.UpdateOrder(data)

		c.Redirect(status.StatusFound, "/orders/"+stringId)
	})

	r.GET("/orders/new", func(c *gin.Context) {
		c.HTML(status.StatusOK, "new.tmpl", gin.H{
			"Proteins": dto.Proteins.GetProteins(),
			"Toppings": dto.Toppings.GetToppings(),
		})
	})

	r.POST("/orders", func(c *gin.Context) {
		data := &dto.Order{}

		if err := c.ShouldBind(data); err != nil {
			c.HTML(status.StatusOK, "new.tmpl", gin.H{
				"Message":  "No la compliques 😞",
				"Proteins": dto.Proteins.GetProteins(),
				"Toppings": dto.Toppings.GetToppings(),
			})
			return
		}

		if data.OrderType == "" {
			c.HTML(status.StatusBadRequest, "new.tmpl", gin.H{
				"Message":  "Elegi el tipo de proteina pa 🐔🐖🌱!",
				"Proteins": dto.Proteins.GetProteins(),
				"Toppings": dto.Toppings.GetToppings(),
			})
			return
		}

		if data.User == "" {
			c.HTML(status.StatusBadRequest, "new.tmpl", gin.H{
				"Message":  "El nombre es mucho muy importante ⚠️ ",
				"Proteins": dto.Proteins.GetProteins(),
				"Toppings": dto.Toppings.GetToppings(),
			})
			return
		}

		r.service.AddOrder(data)

		c.Redirect(status.StatusFound, "/orders")
	})

	r.GET("/orders/:id/delete", func(c *gin.Context) {
		stringId := c.Param("id")
		id, err := strconv.Atoi(stringId)

		if err != nil {
			c.HTML(status.StatusBadRequest, "error.tmpl", gin.H{
				"Message": "Invalid id",
			})
		}

		r.service.DeleteOrder(id)

		c.Redirect(status.StatusFound, "/orders")
	})
}
