package utils

import "reflect"

// CopyNonEmptyFields copies the non-empty fields from src to dest
func CopyNonEmptyFields(dest, src interface{}) {
	// get the reflect.Value of the dest and src
	destValue := reflect.ValueOf(dest).Elem()
	srcValue := reflect.ValueOf(src).Elem()

	for i := 0; i < destValue.NumField(); i++ {
		// get the field of the dest and src
		destField := destValue.Field(i)
		srcField := srcValue.Field(i)

		// check if scrField is of type slice
		if srcField.Kind() == reflect.Slice {
			// if slice is not empty, copy it to destField
			if srcField.Len() > 0 {
				destField.Set(srcField)
			}
		} else {
			// if field is not empty, copy it to destField
			if srcField.Interface() != reflect.Zero(srcField.Type()).Interface() {
				destField.Set(srcField)
			}
		}
	}
}
