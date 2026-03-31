package test

import (
	"reflect"
	"testing"
)
type user struct{
	Userld string
	Nama string
}
func TestReflectPrt(t *testing.T){
	var (
		model *user
		st reflect.Type
		elem reflect.Value
	)
	// model = &user{}
	// sv = reflect.ValueOf(model)
	// t.Log("reflect.ValueOf",sv.Kind().String())
	// sv = sv.Elem()
	// t.Log("reflect.ValueOf.elem",sv.Kind().String())
	// sv.FieldByName("Userld").SetString("123456")
	// sv.FieldByName("Name").SetString("nickname")
	// t.Log("model",model)
	st = reflect.TypeOf(model)
	t.Log("reflect.TypeOf",st.Kind().String())
	st = st.Elem()
	t.Log("reflect.TypeOf.Elem",st.Kind().String())
	elem = reflect.New(st)
	t.Log("reflect.New",elem.Kind().String())
	t.Log("reflect.New.elem",elem.Elem().Kind().String())
	model = elem.Interface().(*user)
	elem= elem.Elem()
	elem.FieldByName("Userld").SetString("12345")
	elem.FieldByName("Name").SetString("nihao")
	t.Log("model model.Name",model,model.Nama)

}
