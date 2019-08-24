package function

import (
	"fmt"
	"reflect"
)

type GenericFunction struct {
	function func(...interface{})
}

type GenericInterface interface {
	FuncName() func(...interface{})
}

func (gf GenericFunction) blahblahblah() {

}

func (gf GenericFunction) FuncName(name string) func(...interface{}) {

	doo := reflect.ValueOf(&gf).MethodByName(name)

	fmt.Println("hoop")
	fmt.Println(doo.String())
	fmt.Println("DOOP")

	return func(...interface{}){}
}

func (gf GenericFunction) Blah(name string) {


}

func (f *Function) Execute() GenericFunction {
/*
	var gf GenericFunction



	gf.args = f.Args
	reflect.ValueOf(&gi.FuncName()).MethodByName(f.Name)

	fmt.Println(blah)
*/
	return GenericFunction{}
}