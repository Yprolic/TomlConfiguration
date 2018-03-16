package TomlConfiguration

import (
	"fmt"
	"reflect"
	"strconv"
)

type TagLoader struct {
	DefaultTagName string
}

func (t *TagLoader) Load(s interface{}) error {
	if t.DefaultTagName == "" {
		t.DefaultTagName = "default"
	}
	sv := reflect.ValueOf(s).Elem()
	st := sv.Type()
	for i := 0; i < st.NumField(); i++ {
		defValue := st.Field(i).Tag.Get("default")
		if defValue != "" {
			setField(sv.Field(i), defValue)
		}
	}
	return nil
}
func setField(field reflect.Value, v string) error {
	switch field.Kind() {
	case reflect.Bool:
		if val, err := strconv.ParseBool(v); err != nil {
			return err
		} else {
			field.SetBool(val)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if val, err := strconv.ParseInt(v, 0, 64); err != nil {
			return err
		} else {
			field.SetInt(val)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if val, err := strconv.ParseUint(v, 0, 64); err != nil {
			return err
		} else {
			field.SetUint(val)
		}
	case reflect.Float32, reflect.Float64:
		if val, err := strconv.ParseFloat(v, 64); err != nil {
			return err
		} else {
			field.SetFloat(val)
		}
	case reflect.String:
		field.SetString(v)
	default:
		return fmt.Errorf("multiconfig: field '%s' has unsupported type: %s", field.String(), field.Kind())
	}

	return nil
}
