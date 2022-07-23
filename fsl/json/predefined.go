package json

import (
	"errors"
	"fmt"

	"github.com/vyeve/fsl/fsl"
)

func (p *parser) create(cmd, params command) error {
	variable, err := p.getVariableNameByID(idKey, cmd, nil) // extract name of variable
	if err != nil {
		return err
	}
	_, err = p.getValueByID(variable) // validate if variable not exist
	if err == nil {
		return fmt.Errorf("%w: %s", fsl.ErrVariableAlreadyExist, variable)
	}
	value, err := p.extractValue(valueKey, cmd, params)
	if err != nil {
		return err
	}
	p.setValue(variable, value)
	return nil
}

func (p *parser) delete(cmd command) error {
	variable, err := p.getVariableNameByID(idKey, cmd, nil) // extract name of variable
	if err != nil {
		return err
	}
	delete(p.variables, variable)
	return nil
}

func (p *parser) update(cmd, params command) error {
	variable, err := p.getVariableNameByID(idKey, cmd, nil) // extract name of variable
	if err != nil {
		return err
	}
	_, err = p.getValueByID(variable) // validate if variable not exist
	if err != nil {
		return err
	}
	value, err := p.extractValue(valueKey, cmd, params)
	if err != nil {
		return err
	}
	p.setValue(variable, value)
	return nil
}

func (p *parser) add(cmd, params command) error {
	return p.calculate(cmd, params, func(value1, value2 float64) float64 {
		return value1 + value2
	})
}

func (p *parser) subtract(cmd, params command) error {
	return p.calculate(cmd, params, func(value1, value2 float64) float64 {
		return value1 - value2
	})
}

func (p *parser) multiply(cmd, params command) error {
	return p.calculate(cmd, params, func(value1, value2 float64) float64 {
		return value1 * value2
	})
}

func (p *parser) divide(cmd, params command) error {
	return p.calculate(cmd, params, func(value1, value2 float64) float64 {
		return value1 / value2
	})
}

func (p *parser) calculate(cmd, params command, fn func(value1, value2 float64) float64) error {
	varName, err := p.getVariableNameByID(idKey, cmd, params) // extract name of variable
	if err != nil {
		return err
	}
	value1, err := p.extractValue(operand1Key, cmd, params)
	if err != nil {
		return err
	}
	value2, err := p.extractValue(operand2Key, cmd, params)
	if err != nil {
		return err
	}
	p.setValue(varName, fn(value1, value2))
	return nil
}

func (p *parser) print(cmd, params command) error {
	value, err := p.extractValue(valueKey, cmd, params)
	switch {
	case err == nil:
		_, err = fmt.Fprintln(p.writer, value)
	case errors.Is(err, fsl.ErrVarNotFound):
		_, err = fmt.Fprintln(p.writer, undefinedKeyword)
	default:
		return err
	}
	return err
}
