package main

import (
	"fmt"

	"github.com/go-summer/cmd/base-example/datagenerator"
	"github.com/go-summer/cmd/base-example/writer"
)

func main() {
	fmt.Println("-----start-----")

	writer.NewPebble("test1")

	dataGenerator := datagenerator.NewDataGenerator()
	dataGenerator.Generate()

	fmt.Println("-----finish-----")
}
