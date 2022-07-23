package json

import (
	"encoding/json"
	"fmt"

	"github.com/vyeve/fsl/fsl"
)

// init tries to unmarshal bytes to map[string]interface{} and set variables and functions
func (p *parser) init(data []byte) error {
	m := make(map[string]interface{})
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch v := v.(type) {
		case float64:
			p.variables[k] = v
		case []interface{}:
			cmds := make([]command, len(v))
			for i, c := range v {
				cmd, ok := c.(command)
				if !ok {
					return fmt.Errorf("%w: expected map[string]interface{}, got: %T", fsl.ErrIncorrectInputData, c)
				}
				cmds[i] = cmd
			}
			p.functions[k] = cmds
		default:
			return fmt.Errorf("%w: unexpected type %T", fsl.ErrIncorrectInputData, v)
		}
	}
	return nil
}

// parseFunction receives name of function and optional params
func (p *parser) parseFunction(name string, params command) error {
	fn, ok := p.functions[name]
	if !ok {
		return fmt.Errorf("%w: %s", fsl.ErrFunctionNotFound, name)
	}
	for _, cmd := range fn {
		command, err := p.getVariableNameByID(cmdKey, cmd, nil)
		if err != nil {
			return err
		}
		switch command {
		case createFnKey:
			err = p.create(cmd, params)
		case deleteFnKey:
			err = p.delete(cmd)
		case updateFnKey:
			err = p.update(cmd, params)
		case addFnKey:
			err = p.add(cmd, params)
		case subtractFnKey:
			err = p.subtract(cmd, params)
		case multiplyFnKey:
			err = p.multiply(cmd, params)
		case divideFnKey:
			err = p.divide(cmd, params)
		case printFnKey:
			err = p.print(cmd, params)
		default:
			// not predefined function
			err = p.parseFunction(command, cmd)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
