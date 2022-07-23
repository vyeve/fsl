package json

import (
	"fmt"

	"github.com/vyeve/fsl/fsl"
)

func (p *parser) getVariableNameByID(id string, cmd, params command) (string, error) {
	v, ok := cmd[id]
	if !ok {
		return "", fmt.Errorf("%w: %s", fsl.ErrVarNotFound, id)
	}
	id, ok = v.(string)
	if !ok {
		return "", fmt.Errorf("%w: %+v", fsl.ErrIncorrectInputData, v)
	}

	if len(id) == 0 {
		return "", fmt.Errorf("%w: %s", fsl.ErrIncorrectInputData, v)
	}
	switch id[0] {
	case '#':
		return id[1:], nil
	case '$':
		return p.getVariableNameByID(id[1:], params, nil)
	default:
		return id, nil
	}
}

func (p *parser) extractValue(id string, cmd, params command) (float64, error) {
	v, ok := cmd[id]
	if !ok {
		return 0, fmt.Errorf("%w: %s", fsl.ErrVarNotFound, id)
	}
	switch v := v.(type) {
	case float64:
		return v, nil
	case string:
		id = v
	default:
		return 0, fmt.Errorf("%w: %s [%T]", fsl.ErrIncorrectInputData, id, v)
	}
	if len(id) == 0 {
		return 0, fsl.ErrIncorrectInputData
	}
	switch id[0] {
	case '#': // reference to variable
		id = id[1:]
	case '$': // argument to the function
		return p.extractValue(id[1:], params, nil)
	default:
		return 0, fmt.Errorf("%w: %s", fsl.ErrIncorrectInputData, id)
	}
	return p.getValueByID(id)
}

func (p *parser) getValueByID(id string) (float64, error) {
	value, ok := p.variables[id]
	if !ok {
		return 0, fmt.Errorf("%w: %s", fsl.ErrVarNotFound, id)
	}
	return value, nil
}

func (p *parser) setValue(id string, value float64) {
	p.variables[id] = value
}
