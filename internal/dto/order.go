package dto

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

type Order struct {
	ID        int           `json:"id" form:"id" validate:"omitempty" emsg:"ID invalido"`
	OrderType ShawarmaType  `json:"orderType" form:"orderType" validate:"required,validate_shawarma" emsg:"Elegi el tipo de proteina pa 🐔🐖🌱!"`
	Toppings  []ToppingType `json:"toppings" form:"toppings" validate:"omitempty,validate_toppings" emsg:"Toppings invalidos"`
	User      string        `json:"user" form:"user" validate:"required" emsg:"El nombre es mucho muy importante ⚠️"`
	IsActive  bool          `json:"isActive" form:"isActive"`
}

func ValidateToppings(fl validator.FieldLevel) bool {
	if fl.Field().Kind() == reflect.Slice {
		toppings := fl.Field().Interface().([]ToppingType)
		for _, topping := range toppings {
			if _, err := GetTopping(topping); err != nil {
				return false
			}
		}
		return true
	}
	return false
}

func ValidateShawarma(fl validator.FieldLevel) bool {
	if fl.Field().Kind() == reflect.String {
		shawarmaType := fl.Field().Interface().(ShawarmaType)
		if _, err := GetShawarma(shawarmaType); err != nil {
			return false
		}
		return true
	}
	return false
}

func (m *Order) Validate(validate *validator.Validate) error {
	validate.RegisterValidation("validate_toppings", ValidateToppings)
	validate.RegisterValidation("validate_shawarma", ValidateShawarma)
	return validateFunc[Order](*m, validate)
}
