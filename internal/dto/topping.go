package dto

import "fmt"

type ToppingKey string

const (
	HumusValue                 ToppingKey = "154249"
	TomateValue                ToppingKey = "154250"
	MayonesaValue              ToppingKey = "154251"
	RepolloConZanahoriaValue   ToppingKey = "154252"
	PicanteValue               ToppingKey = "154253"
	PepinoValue                ToppingKey = "160336"
	CebollaValue               ToppingKey = "154254"
	SalsaDeTomateConCurryValue ToppingKey = "154255"
	YogurtValue                ToppingKey = "154256"
	SalsaDeSesamoThinaValue    ToppingKey = "154257"
	PapasPayValue              ToppingKey = "154258"
)

type ToppingValue string

const (
	HumusKey                 ToppingValue = "Humus"
	TomateKey                ToppingValue = "Tomate"
	MayonesaKey              ToppingValue = "Mayonesa"
	RepolloConZanahoriaKey   ToppingValue = "Repollo con Zanahoria"
	PicanteKey               ToppingValue = "Picante"
	PepinoKey                ToppingValue = "Pepino"
	CebollaKey               ToppingValue = "Cebolla"
	SalsaDeTomateConCurryKey ToppingValue = "Salsa de Tomate con Curry"
	YogurtKey                ToppingValue = "Yogurt"
	SalsaDeSesamoThinaKey    ToppingValue = "Salsa de Sesamo Thina"
	PapasPayKey              ToppingValue = "Papas Pay"
)

type Topping struct {
	ID   ToppingKey   `json:"id"`
	Name ToppingValue `json:"name"`
}

type ToppingsMap map[ToppingKey]ToppingValue

var Toppings = ToppingsMap{
	HumusValue:                 HumusKey,
	TomateValue:                TomateKey,
	MayonesaValue:              MayonesaKey,
	RepolloConZanahoriaValue:   RepolloConZanahoriaKey,
	PicanteValue:               PicanteKey,
	PepinoValue:                PepinoKey,
	CebollaValue:               CebollaKey,
	SalsaDeTomateConCurryValue: SalsaDeTomateConCurryKey,
	YogurtValue:                YogurtKey,
	SalsaDeSesamoThinaValue:    SalsaDeSesamoThinaKey,
	PapasPayValue:              PapasPayKey,
}

func (t ToppingsMap) GetToppingName(id ToppingKey) (ToppingValue, error) {
	if val, ok := t[id]; ok {
		return val, nil
	}
	return "", fmt.Errorf("topping not found")
}

func (t ToppingsMap) GetToppings() []Topping {
	var toppings []Topping
	for k, v := range t {
		toppings = append(toppings, Topping{ID: k, Name: v})
	}
	return toppings
}
