package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/vyeve/fsl/fsl/json"
)

type array []string

func (a *array) Set(in string) error {
	*a = append(*a, in)
	return nil
}

func (a array) String() string {
	return strings.Join(a, " ")
}

func main() {
	var files array
	flag.Var(&files, "f", "path to FSL file")
	flag.Parse()

	p := json.NewParser(os.Stdout)
	for _, f := range files {
		b, err := ioutil.ReadFile(f)
		if err != nil {
			log.Fatal(err)
		}
		err = p.Parse(b)
		if err != nil {
			log.Fatal(err)
		}
	}
}
