package main
import(
	"fmt"
)
func main(){
	a1 := [...]int{1,3,5,7,8}

	m := make(map[int]int)

	for i, v := range a1 {
    	target := 8 - v
		/*
		if idx, ok := m[target]; ok {
	    fmt.Println(idx, i)
		}
		它等价于三步：
		第一步：从 map 里取值
		idx := m[target]
		但问题是：map 里可能没有这个 key。
		第二步：判断这个 key 在不在 map 里
		Go 提供了一个“安全取值方式”：
		value, ok := map[key]
		value：取到的值
		ok：布尔值
		true 👉 key 存在
		false 👉 key 不存在
		第三步：组合成 if
		if idx, ok := m[target]; ok



		*/

    	if idx, ok := m[target]; ok { //根据map key查vlue 并赋值
        	fmt.Println(idx, i)
    	}

    	m[v] = i
	}
}