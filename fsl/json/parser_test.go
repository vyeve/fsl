package json

import (
	"io/ioutil"
	"log"
	"os"
)

func ExampleParser() {
	data, err := ioutil.ReadFile("../../sample/sample1.json")
	if err != nil {
		log.Fatal(err)
	}
	p := NewParser(os.Stdout)
	err = p.Parse(data)
	if err != nil {
		log.Fatal(err)
	}
	// Output:
	// 3.5
	// 5.5
	// 5
	// 25
	// undefined
	// 2
	// 5
}
