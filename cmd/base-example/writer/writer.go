package writer

import (
	"fmt"

	"github.com/go-summer/internal/core/pebble"
	"github.com/go-summer/internal/core/sontext"
)

type Writer interface {
	Write(message string)
}

type writerImpl struct {
	prefix string
	p      pebble.Metadata
}

func NewPebble(prefix string) {
	w := writerImpl{prefix: prefix}
	var wInterface interface{} = &w
	w.p = pebble.NewMetadata(&wInterface, pebble.TypeOf[Writer]())
	sontext.Register(&w)
}

func (w writerImpl) Metadata() pebble.Metadata {
	return w.p
}

func (w writerImpl) Write(message string) {
	fmt.Printf("%s: %s\n", w.prefix, message)
}
