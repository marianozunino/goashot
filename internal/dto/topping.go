package dto

import (
	"fmt"
	"sort"
)

type ToppingType string

const (
	ToppingTypeHumus                 ToppingType = "154249"
	ToppingTypeTomate                            = "154250"
	ToppingTypeMayonesa                          = "154251"
	ToppingTypeRepolloConZanahoria               = "154252"
	ToppingTypePicante                           = "154253"
	ToppingTypePepino                            = "160336"
	ToppingTypeCebolla                           = "154254"
	ToppingTypeSalsaDeTomateConCurry             = "154255"
	ToppingTypeYogurt                            = "154256"
	ToppingTypeSalsaDeSesamoThina                = "154257"
	ToppingTypePapasPay                          = "154258"
)

func (t ToppingType) String() string {
	return string(t)
}

type Topping struct {
	Name string      `json:"name"`
	ID   ToppingType `json:"id"`
}

var Toppings = map[ToppingType]Topping{
	ToppingTypeHumus:                 {Name: "Humus", ID: ToppingTypeHumus},
	ToppingTypeTomate:                {Name: "Tomate", ID: ToppingTypeTomate},
	ToppingTypeMayonesa:              {Name: "Mayonesa", ID: ToppingTypeMayonesa},
	ToppingTypeRepolloConZanahoria:   {Name: "Repollo con zanahoria", ID: ToppingTypeRepolloConZanahoria},
	ToppingTypePicante:               {Name: "Picante", ID: ToppingTypePicante},
	ToppingTypePepino:                {Name: "Pepino", ID: ToppingTypePepino},
	ToppingTypeCebolla:               {Name: "Cebolla", ID: ToppingTypeCebolla},
	ToppingTypeSalsaDeTomateConCurry: {Name: "Salsa de tomate con curry", ID: ToppingTypeSalsaDeTomateConCurry},
	ToppingTypeYogurt:                {Name: "Yogurt", ID: ToppingTypeYogurt},
	ToppingTypeSalsaDeSesamoThina:    {Name: "Salsa de sesamo thina", ID: ToppingTypeSalsaDeSesamoThina},
	ToppingTypePapasPay:              {Name: "Papas pay", ID: ToppingTypePapasPay},
}

func GetTopping(toppingType ToppingType) (Topping, error) {
	if topping, ok := Toppings[toppingType]; ok {
		return topping, nil
	}
	return Topping{}, fmt.Errorf("topping not found: %s", toppingType)
}

func GetAllToppings() []Topping {
	var result []Topping
	for _, topping := range Toppings {
		result = append(result, topping)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result
}
