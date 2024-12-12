package main
import (
	"fmt"
	"strconv"

	"time"
)

func test(){
	for i :=1 ; i <= 50 ;i++ {
		fmt.Println("hello golang" + strconv.Itoa(i))
		//阻塞1秒
		time.Sleep(time.Second)
	}
}
func task(name string) {
    for i := 0; i < 5; i++ {
        fmt.Printf("%s is running: %d\n", name, i)
        time.Sleep(100 * time.Millisecond)
    }
}
func main(){	
	go test()// 开启协程
	go task("Goroutine 1")
	go task("Goroutine 2")
	for i :=1 ; i <= 10 ;i++ {
		fmt.Println("hello cx" + strconv.Itoa(i))
		//阻塞1秒
		time.Sleep(time.Second)
	}

}
