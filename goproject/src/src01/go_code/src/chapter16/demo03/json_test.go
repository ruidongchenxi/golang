package main
import (
	"testing"
	//"fmt"
)
func TestWriter(t *testing.T) {
	Writer()
	
}
func TestRead(t *testing.T) {
	
	Read()
}
func TestStore(t *testing.T){
	//创建
	monster :=Monster{
		Nama:"牛魔王",
		Age: 5000,
		Skill: "牛魔拳",
	}
	res:=monster.Store()
	if !res {
		t.Fatalf("序列化失败")
	}
	t.Logf("测试成功")
}
func TestReStore(t *testing.T){
	//创建
	monster :=Monster{}
	res:=monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore()错误，希望是=%v，实际=%v",true,res)
	}
	if monster.Nama !="牛魔王"{
		t.Fatalf("monster.ReStore()错误，希望=%v。实际=%v","牛魔王",monster.Nama)
	}
	t.Logf("测试成功")
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src\chapter16\demo03> go test -v -run TestStore
// === RUN   TestStore
//     json_test.go:25: 测试成功
// --- PASS: TestStore (0.00s)
// PASS
// ok      src01/go_code/src/chapter16/demo03      0.464s
// PS D:\golang\goproject\src\src01\go_code\src\chapter16\demo03> go test -v
// === RUN   TestWriter
// --- PASS: TestWriter (0.00s)
// === RUN   TestRead
// {孙悟空 1500 七十二变}
// --- PASS: TestRead (0.00s)
// === RUN   TestStore
//     json_test.go:25: 测试成功
// --- PASS: TestStore (0.00s)
// === RUN   TestReStore
//     json_test.go:37: 测试成功
// --- PASS: TestReStore (0.00s)
// PASS
// ok      src01/go_code/src/chapter16/demo03      0.410s