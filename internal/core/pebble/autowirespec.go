package pebble

import "reflect"

type AutowireSpec interface {
	Name() string
	Type() reflect.Type
	Inject(pebble Pebble)
}

type autowireSpec struct {
	name       string
	ttype      reflect.Type
	fieldValue reflect.Value
}

func (receiver *autowireSpec) Name() string {
	return receiver.name
}

func (receiver *autowireSpec) Type() reflect.Type {
	return receiver.ttype
}

func (receiver *autowireSpec) Inject(pebble Pebble) {
	pebbleValue := reflect.ValueOf(pebble)
	fieldType := receiver.fieldValue.Type()

	// Check if pebble implements the required type
	if !pebbleValue.Type().Implements(receiver.ttype) {
		return
	}

	// If field is a pointer, create a pointer to the interface type
	if fieldType.Kind() == reflect.Ptr {
		ptr := reflect.New(receiver.ttype)
		ptr.Elem().Set(pebbleValue)
		receiver.fieldValue.Set(ptr)
	} else {
		// Field is not a pointer, set directly
		receiver.fieldValue.Set(pebbleValue)
	}
}

// NewAutowireSpec creates an AutowireSpec for a field that can be injected.
// fieldPtr should be a pointer to the field you want to inject into.
// Example: NewAutowireSpec(&structInstance.fieldName)
func NewAutowireSpec(fieldPtr interface{}) AutowireSpec {
	fieldValue := reflect.ValueOf(fieldPtr)
	if fieldValue.Kind() != reflect.Ptr {
		panic("NewAutowireSpec: fieldPtr must be a pointer to the field")
	}

	// Get the actual field value (dereference the pointer we passed: &field -> field)
	fieldValue = fieldValue.Elem()

	// Get the type of the field (what type we're looking for to inject)
	ttype := fieldValue.Type()
	// If the field itself is a pointer type (e.g., *writer.Writer), get the element type (writer.Writer)
	if ttype.Kind() == reflect.Ptr {
		ttype = ttype.Elem()
	}

	return &autowireSpec{
		name:       "",
		ttype:      ttype,
		fieldValue: fieldValue,
	}
}
