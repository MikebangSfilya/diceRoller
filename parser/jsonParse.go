package parser

import (
	"encoding/json"
)

type JSONParse struct{}

func (p *JSONParse) Parse(input []byte) ([]PersonRequest, error) {
	persons := []PersonRequest{}
	err := json.Unmarshal(input, &persons)
	if err != nil {
		return []PersonRequest{}, err
	}
	return persons, nil
}
