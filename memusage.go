package memusage

import (
	"reflect"
)

// Gets the memory footprint (in bytes) of the passed parameter
// Please pass the pointer
func GetSize(value interface{}) (int64, error) {
	return getSize(reflect.ValueOf(value))
}

func getSize(rValue reflect.Value) (int64, error) {

	var rv reflect.Value = rValue

	if rValue.Kind() == reflect.Pointer {
		rv = rValue.Elem()
	}

	if !rv.IsValid() {
		if rValue.Kind() == reflect.Pointer {
			return int64(rValue.Type().Size()), nil
		}
		return 0, nil
	}

	var (
		size      int64 = int64(rv.Type().Size())
		fieldSize int64 = 0
		err       error = nil
	)

	switch rv.Kind() {
	case reflect.Map:

		maplist := rv.MapRange()
		for maplist.Next() {
			var keySize int64 = 0
			keySize, err = getSize(maplist.Key())
			if err != nil {
				continue
			}

			fieldSize, err = getSize(maplist.Value())
			if err != nil {
				continue
			}

			size += keySize + fieldSize
		}
	case reflect.Struct:
		size = 0
		for i := 0; i < rv.NumField(); i++ {
			fieldRv := rv.Field(i)
			fieldRv.Type()

			if !rv.IsValid() {
				continue
			}
			fieldSize, err = getSize(fieldRv)
			if err != nil {
				continue
			}

			size += fieldSize

		}
	case reflect.String:
		size += int64(len(rv.String()))
	case reflect.Slice:
		if rv.Cap() > 0 {
			for n := 0; n < rv.Len(); n++ {
				fieldSize, err = getSize(rv.Index(n))
				if err != nil {
					continue
				}
				size += fieldSize
			}
		}
		// default:
	}

	return size, err
}
