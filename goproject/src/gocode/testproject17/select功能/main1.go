package main
import(
	"fmt"
	_"time"
	"sync"
)
var wg sync.WaitGroup
//hanshu
func prinmun(){
	defer wg.Done()
	for i := 1;i <=10;i++{
		fmt.Println(i)
	//	time.Sleep(time.Second)
	}
}
func devide(){
	defer wg.Done()
	defer func(){
		err := recover()
		if err != nil{
			fmt.Println("cuowu")
		}
	}()
	//defer wg.Done()
	num1 := 0
	num2 := 20
	result := num2 / num1
	//time.Sleep(time.Second)
	fmt.Println(result)
	
}

func main(){
	wg.Add(2)
    go prinmun()
	go devide()
	wg.Wait()
}