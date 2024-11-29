package main
import "fmt"
func main(){
   var sum int 
   for i :=1; i<=100;i++ {    
      sum += i
      fmt.Println(sum)
      if sum >=300 {
         break   //停止正在执行的循环
      }
   }
   ftm.Println("ok")
}