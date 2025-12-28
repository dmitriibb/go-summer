package pebble

import "reflect"

type AutowireSpec interface {
	Name() string
	Type() reflect.Type
	Inject(*interface{})
}

type autowireSpec struct {
	name           string
	ttype          reflect.Type
	injectionField *interface{}
}

func (receiver *autowireSpec) Name() string {
	return receiver.name
}

func (receiver *autowireSpec) Type() reflect.Type {
	return receiver.ttype
}

func (receiver *autowireSpec) Inject(injectionObject *interface{}) {
	receiver.injectionField = injectionObject
}

func NewAutowireSpec(injectionField *interface{}) AutowireSpec {
	ttype := reflect.TypeOf(*injectionField)
	if ttype.Kind() == reflect.Ptr {
		ttype = ttype.Elem()
	}
	return &autowireSpec{
		name:           "",
		ttype:          ttype,
		injectionField: injectionField,
	}
}
