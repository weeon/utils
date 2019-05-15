package utils

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
)

func ExecFuncs(fns []func() error) error {
	for _, fn := range fns {
		err := fn()
		if err != nil {
			return errors.New(fmt.Sprintf("Run %v Error: %v \n", GetFunctionName(fn), err))
		}
	}
	return nil
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
