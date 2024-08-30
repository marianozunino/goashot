package handler

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/marianozunino/goashot/internal/dto"
	"github.com/marianozunino/goashot/internal/service"
	"github.com/marianozunino/goashot/internal/templates"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type OrdersRessources struct {
	OrdersService service.OrderService
}

func (rs OrdersRessources) Routes(e *echo.Echo) {
	ordersGroup := e.Group("/orders")

	ordersGroup.GET("", rs.getAllOrders)
	ordersGroup.GET("/:id", rs.getOrders)

	ordersGroup.GET("/new", rs.getNewOrders)
	ordersGroup.POST("", rs.postOrders)
	ordersGroup.PUT("/:id", rs.putOrders)
	ordersGroup.DELETE("/:id", rs.deleteOrders)
}

func (rs OrdersRessources) getAllOrders(e echo.Context) error {
	data := rs.OrdersService.GetOrders()
	return Render(e, http.StatusOK, templates.Index(data))
}

func (rs OrdersRessources) getNewOrders(e echo.Context) error {
	template := templates.OrderPage(templates.OrderContext{
		IsEdit:    false,
		Order:     nil,
		Shawarmas: dto.GetAllShawarmas(),
		Toppings:  dto.GetAllToppings(),
	})

	return Render(e, http.StatusOK, template)

}

func (rs OrdersRessources) postOrders(e echo.Context) error {
	data := &dto.Order{}

	if err := e.Bind(data); err != nil {
		return err
	}

	if err := data.Validate(validate); err != nil {
		return Render(e, http.StatusUnprocessableEntity, templates.Alert(templates.AlertTypeError, err.Error()))
	}

	order, err := rs.OrdersService.AddOrder(data)
	if err != nil {
		return Render(e, http.StatusUnprocessableEntity, templates.Alert(templates.AlertTypeError, err.Error()))
	}

	e.Set("HX-Location", "/orders/"+strconv.Itoa(order.ID))
	return e.Redirect(http.StatusSeeOther, "/orders/"+strconv.Itoa(order.ID))
}

func (rs OrdersRessources) getOrders(e echo.Context) error {
	id := e.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return Render(e, http.StatusUnprocessableEntity, templates.Alert(templates.AlertTypeError, "Invalid id"))
	}

	order, err := rs.OrdersService.GetOrder(idInt)
	if err != nil {
		return Render(e, http.StatusUnprocessableEntity, templates.Alert(templates.AlertTypeError, "Order not found"))
	}

	return Render(e, http.StatusOK, templates.OrderPage(templates.OrderContext{
		IsEdit:    true,
		Order:     order,
		Shawarmas: dto.GetAllShawarmas(),
		Toppings:  dto.GetAllToppings(),
	}))
}

func (rs OrdersRessources) putOrders(e echo.Context) error {
	data := &dto.Order{}

	if err := e.Bind(data); err != nil {
		return err
	}

	if err := data.Validate(validate); err != nil {
		return Render(e, http.StatusUnprocessableEntity, templates.Alert(templates.AlertTypeError, err.Error()))
	}

	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return Render(e, http.StatusUnprocessableEntity, templates.Alert(templates.AlertTypeError, "Invalid id"))
	}

	data.ID = id
	updatedOder, err := rs.OrdersService.UpdateOrder(data)

	if err != nil {
		return Render(e, http.StatusUnprocessableEntity, templates.Alert(templates.AlertTypeError, err.Error()))
	}

	return Render(e, http.StatusOK,
		templates.Alert(templates.AlertTypeSuccess, "Actualizado ✔️"),
		templates.OrderPage(templates.OrderContext{
			IsEdit:    true,
			Order:     updatedOder,
			Shawarmas: dto.GetAllShawarmas(),
			Toppings:  dto.GetAllToppings(),
		}),
	)
}

func (rs OrdersRessources) deleteOrders(e echo.Context) error {

	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return Render(e, http.StatusUnprocessableEntity, templates.Alert(templates.AlertTypeError, "Invalid id"))
	}

	rs.OrdersService.DeleteOrder(id)

	if e.QueryParam("redirect") == "true" {
		e.Set("HX-Location", "/orders/")
	}

	return Render(e, http.StatusOK, templates.Alert(templates.AlertTypeSuccess, ""))
}
