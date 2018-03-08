package TomlConfiguration

import (
	"fmt"
	"reflect"
)

type TagLoader struct {
	DefaultTagName string
}

func (t *TagLoader) Load(s interface{}) error {
	if t.DefaultTagName == "" {
		t.DefaultTagName = "default"
	}
	st := reflect.TypeOf(s)
	for i := 0; i < st.NumField(); i++ {
		fmt.Println(st.Field(i).Tag) //将tag输出出来
	}
	return nil
}
