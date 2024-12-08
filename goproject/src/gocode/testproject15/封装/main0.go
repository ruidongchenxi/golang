package main
import (
	"fmt"
    cy "gocode/testproject15/封装/model"
)

func main(){
	//创建结构体实例
	p := cy.NewPerson("小龙")
    p.SetAge(20)
	g := p.GetAge()
	fmt.Println(g)
}