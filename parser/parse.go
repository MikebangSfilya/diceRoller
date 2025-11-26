package parser

import "fmt"

type Parser interface {
	Parse(input []byte) ([]PersonRequest, error)
}

type ParserManager struct {
	parsers map[string]Parser
}

// Парсер для поддержки нескльких способов ввода
func NewParserManager() *ParserManager {
	pm := &ParserManager{
		parsers: make(map[string]Parser),
	}

	pm.Register("application/json", &JSONParse{})
	return pm
}

func (pm *ParserManager) Register(contentType string, p Parser) {
	pm.parsers[contentType] = p
}

func (pm *ParserManager) Parse(contentType string, input []byte) ([]PersonRequest, error) {
	parser, ok := pm.parsers[contentType]
	if !ok {
		return nil, fmt.Errorf("unsupported content type: %s", contentType)
	}
	return parser.Parse(input)
}
