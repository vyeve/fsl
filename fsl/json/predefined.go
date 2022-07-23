package json

import (
	"errors"
	"fmt"

	"github.com/vyeve/fsl/fsl"
)

func (p *parser) create(data, params map[string]interface{}) error {
	variable, err := p.getVariableNameByID(idKey, data, nil) // extract name of variable
	if err != nil {
		return err
	}
	_, err = p.getValueByID(variable) // validate if variable not exist
	if err == nil {
		return fmt.Errorf("%w: %s", fsl.ErrVariableAlreadyExist, variable)
	}
	value, err := p.extractValue(valueKey, data, params)
	if err != nil {
		return err
	}
	p.setValue(variable, value)
	return nil
}

func (p *parser) delete(data map[string]interface{}) error {
	variable, err := p.getVariableNameByID(idKey, data, nil) // extract name of variable
	if err != nil {
		return err
	}
	delete(p.variables, variable)
	return nil
}

func (p *parser) update(data, params map[string]interface{}) error {
	variable, err := p.getVariableNameByID(idKey, data, nil) // extract name of variable
	if err != nil {
		return err
	}
	_, err = p.getValueByID(variable) // validate if variable not exist
	if err != nil {
		return err
	}
	value, err := p.extractValue(valueKey, data, params)
	if err != nil {
		return err
	}
	p.setValue(variable, value)
	return nil
}

func (p *parser) add(data, params map[string]interface{}) error {
	return p.calculate(data, params, func(value1, value2 float64) float64 {
		return value1 + value2
	})
}

func (p *parser) subtract(data, params map[string]interface{}) error {
	return p.calculate(data, params, func(value1, value2 float64) float64 {
		return value1 - value2
	})
}

func (p *parser) multiply(data, params map[string]interface{}) error {
	return p.calculate(data, params, func(value1, value2 float64) float64 {
		return value1 * value2
	})
}

func (p *parser) divide(data, params map[string]interface{}) error {
	return p.calculate(data, params, func(value1, value2 float64) float64 {
		return value1 / value2
	})
}

func (p *parser) calculate(data, params map[string]interface{}, fn func(value1, value2 float64) float64) error {
	varName, err := p.getVariableNameByID(idKey, data, params) // extract name of variable
	if err != nil {
		return err
	}
	value1, err := p.extractValue(operand1Key, data, params)
	if err != nil {
		return err
	}
	value2, err := p.extractValue(operand2Key, data, params)
	if err != nil {
		return err
	}
	p.setValue(varName, fn(value1, value2))
	return nil
}

func (p *parser) print(data, params map[string]interface{}) error {
	value, err := p.extractValue(valueKey, data, params)
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
