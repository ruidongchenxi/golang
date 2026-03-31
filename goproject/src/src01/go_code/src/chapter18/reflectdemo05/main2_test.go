package test

import (
	"reflect"
	"testing"
)
// type user struct{
// 	Userld string
// 	Nama string
// }
func TestReflectstruct(t *testing.T){
	var (
		model *user
		sv reflect.Value
	)
	model = &user{}
	sv = reflect.ValueOf(model)
	t.Log("reflect.ValueOf",sv.Kind().String())
	sv = sv.Elem()
	t.Log("reflect.ValueOf.elem",sv.Kind().String())
	sv.FieldByName("Userld").SetString("123456")
	sv.FieldByName("Name").SetString("nickname")
	t.Log("model",model)
}
