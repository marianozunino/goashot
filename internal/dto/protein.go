package dto

import "fmt"

type ProteinName string

const (
	ShawarmaCerdo   ProteinName = "Cerdo 🐷"
	ShawarmaPollo   ProteinName = "Pollo 🐔"
	ShawarmaFalafel ProteinName = "Falafel 🌱"
)

type ProteinsMap map[ShawarmaType]ProteinName

type Protein struct {
	ID   ShawarmaType `json:"id"`
	Name ProteinName  `json:"name"`
}

func (p ProteinsMap) GetProtein(id ShawarmaType) (ProteinName, error) {
	if val, ok := p[id]; ok {
		return val, nil
	}
	return "", fmt.Errorf("protein not found")
}

var Proteins = ProteinsMap{
	Cerdo:   ShawarmaCerdo,
	Pollo:   ShawarmaPollo,
	Falafel: ShawarmaFalafel,
}

func (p ProteinsMap) GetProteins() []Protein {
	var proteins []Protein
	for k, v := range p {
		proteins = append(proteins, Protein{ID: k, Name: v})
	}
	return proteins
}
