package json

import (
	"errors"
	"testing"

	"github.com/vyeve/fsl/fsl"
)

func TestExtractValueOK(t *testing.T) {
	/*
		correct cases:
		  "value": 3.5
		  "value": "#var1"  --> "var1": 3.5
		  "value": "$value" --> "value": 3.5
		  "value": "$value" --> "value": "#var1" --> "var1": 3.5
	*/
	expectValue := 1.0
	testCases := []struct {
		name   string
		input  map[string]interface{}
		params map[string]interface{}
	}{
		{
			name: "value is a float",
			input: map[string]interface{}{
				"value": 1.0,
			},
		},
		{
			name: "value is a reference to var1",
			input: map[string]interface{}{
				"value": "#var1",
			},
		},
		{
			name: "value is an argument with value 1.0",
			input: map[string]interface{}{
				"value": "$value",
			},
			params: map[string]interface{}{
				"value": 1.0,
			},
		},
		{
			name: "value is an argument with reference to var1",
			input: map[string]interface{}{
				"value": "$value",
			},
			params: map[string]interface{}{
				"value": "#var1",
			},
		},
	}
	p := parser{
		variables: map[string]float64{
			"var1": 1.0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			value, err := p.extractValue("value", tc.input, tc.params, "")
			if err != nil {
				t.Error(err)
				return
			}
			if value != expectValue {
				t.Errorf("expected %f, got: %f", expectValue, value)
			}
		})
	}
}

func TestExtractValueFailed(t *testing.T) {
	/*
		incorrect cases:
		  "value": "3.5"
		  "value": "#var2"  --> "var2" --> ?
		  "operand1": "#var1"
		  "value": true
		  "value": ""
	*/
	testCases := []struct {
		name   string
		input  map[string]interface{}
		params map[string]interface{}
		err    error
	}{
		{
			name: "value isn't a float",
			input: map[string]interface{}{
				"value": "1.0",
			},
			err: fsl.ErrIncorrectInputData,
		},
		{
			name: "value is a reference to unknown var2",
			input: map[string]interface{}{
				"value": "#var2",
			},
			err: fsl.ErrVarNotFound,
		},
		{
			name: "value is not present in command",
			input: map[string]interface{}{
				"operand1": "#var1",
			},
			err: fsl.ErrVarNotFound,
		},
		{
			name: "value is a bool",
			input: map[string]interface{}{
				"value": true,
			},
			err: fsl.ErrIncorrectInputData,
		},
		{
			name: "empty value",
			input: map[string]interface{}{
				"value": "$value",
			},
			params: map[string]interface{}{
				"value": "",
			},
			err: fsl.ErrIncorrectInputData,
		},
	}
	p := parser{
		variables: map[string]float64{
			"var1": 1.0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := p.extractValue("value", tc.input, tc.params, "")
			if err == nil {
				t.Error("expected not <nil> error")
				return
			}
			if !errors.Is(err, tc.err) {
				t.Errorf("expected %v, got: %v", tc.err, err)
			}
		})
	}
}
