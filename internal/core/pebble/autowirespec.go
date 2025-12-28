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
	if !receiver.fieldValue.IsValid() || !receiver.fieldValue.CanSet() {
		return
	}

	pebbleValue := reflect.ValueOf(pebble)
	fieldType := receiver.fieldValue.Type()

	// Handle pointer fields
	if fieldType.Kind() == reflect.Ptr {
		elemType := fieldType.Elem()

		// Check if pebble implements the interface type (for interface fields like *writer.Writer)
		if elemType.Kind() == reflect.Interface {
			// Field is *Interface, pebble implements Interface
			if pebbleValue.Type().Implements(elemType) {
				// Create a new pointer to the interface and set it
				ptr := reflect.New(elemType)
				ptr.Elem().Set(pebbleValue)
				receiver.fieldValue.Set(ptr)
			}
		} else if pebbleValue.Type().AssignableTo(elemType) {
			// Field is *T, pebble is T - create a pointer
			ptr := reflect.New(pebbleValue.Type())
			ptr.Elem().Set(pebbleValue)
			receiver.fieldValue.Set(ptr)
		} else if pebbleValue.Type().AssignableTo(fieldType) {
			// Field is *T, pebble is *T - direct assignment
			receiver.fieldValue.Set(pebbleValue)
		}
	} else {
		// Handle non-pointer fields
		if fieldType.Kind() == reflect.Interface {
			// Field is Interface, pebble implements Interface
			if pebbleValue.Type().Implements(fieldType) {
				receiver.fieldValue.Set(pebbleValue)
			}
		} else if pebbleValue.Type().AssignableTo(fieldType) {
			// Direct assignment if types match
			receiver.fieldValue.Set(pebbleValue)
		}
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

	// Get the actual field value (dereference the pointer)
	fieldValue = fieldValue.Elem()

	// Get the type of the field (what type we're looking for to inject)
	ttype := fieldValue.Type()
	if ttype.Kind() == reflect.Ptr {
		ttype = ttype.Elem()
	}

	return &autowireSpec{
		name:       "",
		ttype:      ttype,
		fieldValue: fieldValue,
	}
}
