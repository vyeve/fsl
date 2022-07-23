package json

import (
	"testing"
)

func TestExtractValue(t *testing.T) {
	expectValue := 1.0
	testCases := []struct {
		name    string
		needErr bool
		err     error
		input   map[string]interface{}
		params  map[string]interface{}
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
			value, err := p.extractValue("value", tc.input, tc.params)
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
