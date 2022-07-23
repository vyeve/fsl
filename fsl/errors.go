package fsl

import (
	"errors"
)

var (
	ErrVarNotFound          = errors.New("variable not found")
	ErrFunctionNotFound     = errors.New("function not found")
	ErrIncorrectInputData   = errors.New("not correct input data")
	ErrVariableAlreadyExist = errors.New("variable already exists")
)
