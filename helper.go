package gorm_crud

import (
	"fmt"
	"reflect"
	"strconv"
)

func Num64(n interface{}) int64 {
	s := fmt.Sprintf("%d", n)
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	} else {
		return i
	}
}

func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

func NormalizeErr(err error) error {
	if IsNil(err) {
		return nil
	}
	return err
}
