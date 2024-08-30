package dto

import "fmt"

type Shawarma struct {
	ID   ShawarmaType `json:"id"`
	Name string       `json:"name"`
}

type ShawarmaType string

const (
	ShawarmaTypeCerdo   ShawarmaType = "114497_383365"
	ShawarmaTypePollo                = "114497_383366"
	ShawarmaTypeFalafel              = "114497_383367"
)

func (t ShawarmaType) String() string {
	return string(t)
}

var Shawarmas = map[ShawarmaType]Shawarma{
	ShawarmaTypeCerdo:   {Name: "Cerdo 🐷", ID: ShawarmaTypeCerdo},
	ShawarmaTypePollo:   {Name: "Pollo 🐔", ID: ShawarmaTypePollo},
	ShawarmaTypeFalafel: {Name: "Falafel 🌱", ID: ShawarmaTypeFalafel},
}

func GetShawarma(shawarmaType ShawarmaType) (Shawarma, error) {
	if shawarma, ok := Shawarmas[shawarmaType]; ok {
		return shawarma, nil
	}
	return Shawarma{}, fmt.Errorf("shawarma not found: %s", shawarmaType)
}

func GetAllShawarmas() []Shawarma {
	var result []Shawarma
	for _, shawarma := range Shawarmas {
		result = append(result, shawarma)
	}
	return result
}
