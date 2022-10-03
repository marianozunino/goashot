package dto

import "fmt"

type ShawarmaType string

const (
	Cerdo   ShawarmaType = "cerdo"
	Pollo   ShawarmaType = "pollo"
	Falafel ShawarmaType = "falafel"
)

type ProductId string

const (
	ChickenProductId ProductId = "114497_383366"
	PorkProductId    ProductId = "114497_383365"
	FalafelProductId ProductId = "114497_383367"
)

type ShawarmasMap map[ShawarmaType]ProductId

var Shawarmas = ShawarmasMap{
	Cerdo:   PorkProductId,
	Pollo:   ChickenProductId,
	Falafel: FalafelProductId,
}

func (s ShawarmasMap) GetShawarma(id ShawarmaType) (ProductId, error) {
	if val, ok := s[id]; ok {
		return val, nil
	}
	return "", fmt.Errorf("shawarma not found")
}
