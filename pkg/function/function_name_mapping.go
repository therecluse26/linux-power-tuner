package function

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"github.com/therecluse26/uranium/custom"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"reflect"
)

type Funcs map[string]interface{}

var BuiltInFuncs = Funcs{}
var UserFuncs = Funcs{}

func LoadFunctions() {

	UserFuncs["Multiply"] = custom.Multiply
	UserFuncs["UpperCase"] = custom.UpperCase

	fmt.Println("MAIN DIRECTORY = " + viper.Get("ProjectBase").(string))
/*
	err := filepath.Walk(viper.Get("ProjectBase").(string) + "/custom/", UserFuncs.parseFuncs)
	if err != nil {
		panic(err)
	}

	err = filepath.Walk(viper.Get("ProjectBase").(string) + "/custom/", BuiltInFuncs.parseFuncs)
	if err != nil {
		panic(err)
	}
*/
}


func (f Funcs) parseFuncs(path string, info os.FileInfo, _ error) error {

	var err error

	ff := token.NewFileSet()

	if !info.IsDir() {

		source, er := parser.ParseFile(ff, path, nil, 0)
		err = er

		/*for fn := range source.Scope.Objects {
			//f[fn] = source.Scope.Lookup(fn).Decl.(func())
			spew.Dump(source.Scope.Lookup(fn).Decl)
		}*/

		for _, x := range source.Decls {
			fn, ok := x.(*ast.FuncDecl)
			if ok {
				var exported string
				if fn.Name.IsExported() {
					exported = "exported "
				}
				fmt.Println(fn.Name.Obj.Decl)
				fmt.Printf("%sfunction declaration found on line %d: \n\t%s\n", exported, ff.Position(fn.Pos()).Line, fn.Name.Name)
				f[fn.Name.String()] = fn.Name.Obj.Decl
			}
		}

		//spew.Dump(f)

	}
	return err
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

