package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)
func TestAdd(t *testing.T) {
	var dataset = []struct{
		a int
		b int
		out int
	}{
		{1,2,3},
		{12,12,24},
		{-9,8,-1},
		{0,0,0},
	}
	for _,v:= range dataset{
		re:=add(v.a,v.b)
		if re == v.out{
			fmt.Println("正确")
		}
	}
}

func BenchmarkAdd(bb *testing.B) {
	a:=123
	b:=456
	c:=579
	for i:=0;i<bb.N;i++{
		if actual:=add(a,b);actual!=c{
			fmt.Println("aaa ")
		}
	}
}
const num =10000
func BenchmarkStringsprintf(b *testing.B) {
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		var str string
		for j:=0;j<num;j++{
			str=fmt.Sprintf("%s%d",str,j)
		}

	}
	b.StopTimer()
	
}
func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		var str string
		for j:=0;j<num;j++{
			str=str+strconv.Itoa(j)
		}

	}
	b.StopTimer()
	
}
func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		//var str string
		var build strings.Builder
		for j:=0;j<num;j++{
			build.WriteString(strconv.Itoa(j))
			//str=str+strconv.Itoa(j)
		}
		_ = build.String()

	}
	b.StopTimer()
}