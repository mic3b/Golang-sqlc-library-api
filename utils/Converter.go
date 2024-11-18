package utils

import (
	"database/sql"
	"reflect"
	"time"
)

// Convert converts a source object (mainly from Database) to a destination object (mainly from Domain) using reflection
func Convert(src interface{}, dst interface{}) {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst)

	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}
	if dstVal.Kind() == reflect.Ptr {
		dstVal = dstVal.Elem()
	}

	if srcVal.Kind() == reflect.Slice && dstVal.Kind() == reflect.Slice {
		for i := 0; i < srcVal.Len(); i++ {
			srcElem := srcVal.Index(i).Addr().Interface()
			dstElem := reflect.New(dstVal.Type().Elem()).Interface()
			Convert(srcElem, dstElem)
			dstVal.Set(reflect.Append(dstVal, reflect.ValueOf(dstElem).Elem()))
		}
		return
	}

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		dstField := dstVal.FieldByName(srcVal.Type().Field(i).Name)

		if !dstField.IsValid() || !dstField.CanSet() {
			continue
		}

		switch srcField.Interface().(type) {
		case sql.NullTime:
			if srcField.Interface().(sql.NullTime).Valid {
				dstField.Set(reflect.ValueOf(srcField.Interface().(sql.NullTime).Time))
			} else {
				dstField.Set(reflect.ValueOf(time.Time{}))
			}
		case sql.NullInt32:
			if srcField.Interface().(sql.NullInt32).Valid {
				dstField.Set(reflect.ValueOf(srcField.Interface().(sql.NullInt32).Int32))
			} else {
				dstField.Set(reflect.ValueOf(int32(0)))
			}
		case sql.NullString:
			if srcField.Interface().(sql.NullString).Valid {
				dstField.Set(reflect.ValueOf(srcField.Interface().(sql.NullString).String))
			} else {
				dstField.Set(reflect.ValueOf(""))
			}
		case sql.NullFloat64:
			if srcField.Interface().(sql.NullFloat64).Valid {
				dstField.Set(reflect.ValueOf(srcField.Interface().(sql.NullFloat64).Float64))
			} else {
				dstField.Set(reflect.ValueOf(float64(0)))
			}
		case sql.NullBool:
			if srcField.Interface().(sql.NullBool).Valid {
				dstField.Set(reflect.ValueOf(srcField.Interface().(sql.NullBool).Bool))
			} else {
				dstField.Set(reflect.ValueOf(false))
			}
		case sql.NullInt64:
			if srcField.Interface().(sql.NullInt64).Valid {
				dstField.Set(reflect.ValueOf(srcField.Interface().(sql.NullInt64).Int64))
			} else {
				dstField.Set(reflect.ValueOf(int64(0)))
			}
		default:
			dstField.Set(srcField)
		}
	}
}
