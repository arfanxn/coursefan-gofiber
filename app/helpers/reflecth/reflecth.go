package reflecth

import "reflect"

// GetTypeName get name of type of the variable
func GetTypeName(myvar any) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

// GetStructFieldTag returns field tag
func GetStructFieldTag(sf reflect.StructField, tagNames ...string) string {
	if len(tagNames) == 0 {
		return string(sf.Tag)
	}
	return sf.Tag.Get(tagNames[0])
}

// IsTypeOf checks if the type of first variable is a type of the second variable, returns true if type equals the first variable and false otherwise.
func IsTypeOf(first, second any) bool {
	return reflect.TypeOf(first) == reflect.TypeOf(second)
}
