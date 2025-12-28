package pebble

import "reflect"

type Metadata interface {
	Name() string
	Types() []reflect.Type
	IsReady() bool
	Ready()
}

type pebbleMetadata struct {
	name   string
	ttypes []reflect.Type
	ready  bool
}

func (p *pebbleMetadata) Name() string {
	return p.name
}

func (p *pebbleMetadata) Types() []reflect.Type {
	return p.ttypes
}

func (p *pebbleMetadata) IsReady() bool {
	return p.ready
}

func (p *pebbleMetadata) Ready() {
	p.ready = true
}

type Pebble interface {
	Metadata() Metadata
}
