package main

import (
	"fmt"

	"github.com/SriVyshnaviPushadapu/UUID_Generator/generator"
)

func main() {
	fmt.Println("Generating UUID using google UUID package")
	uuid := generator.GenerateUUID()
	fmt.Println("Generated UUID:", uuid)
}