package main
import "fmt"
func main(){
   //双层循环
   /*
    for i := 1;i<=5; i++{
      for j :=2 ; j <=4 ; j++{
         if i == 2 && j == 2{
            break
         } 
         fmt.Printf("i:%v,j:%v \n",i,j)
      }

    }
    
    lable1:
    for i := 1;i<=5; i++{
      for j :=2 ; j <=4 ; j++{
         if i == 2 && j == 2{
            break lable1
         } 
         fmt.Printf("i:%v,j:%v \n",i,j)
      }
    }
    */
    lable1:
    for i := 1;i<=5; i++{
      for j :=2 ; j <=4 ; j++{
         if i == 2 && j == 2{
            continue lable1
         } 
         fmt.Printf("i:%v,j:%v \n",i,j)
      }
    } 
   fmt.Println("ok")
}