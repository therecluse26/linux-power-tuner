package function

import (
	"errors"
	"github.com/therecluse26/uranium/custom"
	"github.com/therecluse26/uranium/pkg/builtin"
	"reflect"
)

type Funcs map[string]interface{}

var BuiltInFuncs = Funcs{}
var UserFuncs = Funcs{}

func LoadFunctions() {
	BuiltInFuncs = builtin.ExportedFuncs
	UserFuncs = custom.ExportedFuncs
}



func CallFunction(m map[string]interface{}, name string, params ... interface{}) (result []reflect.Value, err error) {
	var f = reflect.ValueOf(m[name])

	if len(params) != f.Type().NumIn() {
		err = errors.New("Wrong number of arguments supplied to " + name)
		return
	}
	p := make([]reflect.Value, len(params))
	for k, param := range params {
		p[k] = reflect.ValueOf(param)
	}
	result = f.Call(p)
	return result, err
}

