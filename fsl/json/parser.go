package json

import (
	"io"

	"github.com/vyeve/fsl/fsl"
)

type parser struct {
	writer    io.Writer
	variables map[string]float64
	functions map[string][]command
}

type command = map[string]interface{}

var _ fsl.Parser = (*parser)(nil) // make sure that parser implements fsl.Parser interface

func NewParser(wr io.Writer) *parser {
	p := &parser{
		writer:    wr,
		variables: make(map[string]float64),   // variables is a map key-value
		functions: make(map[string][]command), // function is a slice of commands
	}
	return p
}

func (p *parser) Parse(data []byte) error {
	err := p.init(data)
	if err != nil {
		return err
	}
	return p.parseFunction(initFnKey, nil)
}
