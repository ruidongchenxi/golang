package test

import (
	//"fmt"
	//"fmt"
	"reflect"
	"testing"
)
func TestReflectFunc(t *testing.T){
	call1 := func (v1 int,v2 int)  {
		t.Log(v1,v2)
	}
	call2 := func (v1 int,v2 int,s string)  {
		t.Log(v1,v2,s)
	}
	var (
		function reflect.Value
		inValue []reflect.Value
		//n int 
	)
	bridge := func (call interface{},ages...interface{})  {
		n := len(ages)
		inValue = make([]reflect.Value, n)
		for i:=0; i<n;i++{
			inValue[i]=reflect.ValueOf(ages[i])
		}
		function =  reflect.ValueOf(call)
		function.Call(inValue)
	}
	bridge(call1,1,2)
	bridge(call2,1,2,"test")
}