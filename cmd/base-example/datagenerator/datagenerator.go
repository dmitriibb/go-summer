package datagenerator

import (
	"fmt"

	"github.com/go-summer/cmd/base-example/writer"
	"github.com/go-summer/internal/core/pebble"
	"github.com/go-summer/internal/core/sontext"
)

type DataGenerator interface {
	Generate()
}

type dataGenerator struct {
	output writer.Writer
	p      pebble.Metadata
}

func NewDataGenerator() DataGenerator {
	dg := &dataGenerator{}
	var dgInterface interface{} = dg
	dg.p = pebble.NewMetadata(&dgInterface, pebble.TypeOf[DataGenerator]())

	sontext.Build(dg, pebble.NewAutowireSpec(&dg.output))

	return dg
}

func (dg *dataGenerator) Metadata() pebble.Metadata {
	return dg.p
}

func (d *dataGenerator) Generate() {
	for i := 0; i < 10; i++ {
		d.output.Write(fmt.Sprintf("message - %v", i))
	}
}
