package fsl

type Parser interface {
	Parse([]byte) error
}
