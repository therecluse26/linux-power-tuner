package function

import (
	"errors"
	"github.com/therecluse26/uranium/custom"
	"github.com/therecluse26/uranium/pkg/builtin"
	"github.com/therecluse26/uranium/pkg/types"
	"reflect"
)


func LoadFunctions() types.Funcs {
	var AllFuncs = types.Funcs{}
	for k, f := range builtin.ExportedFuncs {
		AllFuncs[k] = f
	}
	for k, f := range custom.ExportedFuncs {
		AllFuncs[k] = f
	}
	return AllFuncs
}

/*
 * The secret sauce that allows for functions to be
 * called dynamically from presets and config files
 */
func CallFunction(m map[string]interface{}, funcData types.Function) (result interface{}, err error) {

	f := reflect.ValueOf(m[funcData.Name])
	if !f.IsValid() {
		err = errors.New("unable to parse function " + funcData.Name)
		return
	}

	if len(funcData.Args) != f.Type().NumIn() {
		err = errors.New("wrong number of arguments passed to " + funcData.Name)
		return
	}

	p := make([]reflect.Value, len(funcData.Args))
	for _, param := range funcData.Args {
		p[param.Order] = reflect.ValueOf(param.Value)
	}

	// Function must return a value of some kind
	result = f.Call(p)[0]
	return result, err
}

