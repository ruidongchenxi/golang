package main

import (
	"fmt"
	"time"
) 

func main(){
    now := time.Now() //当前时间
    // fmt.Printf("type=%T val=%v\n",now,now)// 打印时间的类型，与当前时间
    // fmt.Println("当前年=",now.Year())
    // fmt.Println("当前月=",now.Month())
    // // 可以将返回的month,转成对应的数值
    // fmt.Println("当前月=",int(now.Month()))
    // fmt.Println("当前日=",now.Day())
    // fmt.Println("当前时=",now.Hour())
    // fmt.Println("当前分=",now.Minute())
    // fmt.Println("当前秒=",now.Second())
    // fmt.Printf("当前年月日 %02d-%02d-%02d %02d:%02d:%02d\n",now.Year(),now.Month(),
    // now.Day(),now.Hour(),
    // now.Minute(),now.Second())
    // fmt.Printf(now.Format("2006-01-02 15:04:05"))
    // fmt.Println()
    // fmt.Printf(now.Format("2006-01-02"))
    // fmt.Println()
    // fmt.Printf(now.Format("15:04:05"))
    // fmt.Println()
    // //func Sleeo(d Duration)  
    // time.Sleep(100*time.Millisecond)
    // i :=0
    // for {
    //     i++
    //     //休眠
    //     time.Sleep(time.Second) //1秒
    //     fmt.Println(i)
    //     if i == 3{
    //         break
    //     }
    // }
    // for i:=1;i<=10;i++{
    //     time.Sleep(time.Millisecond * 100) //10 分之一毫秒
    //     fmt.Println(i)
    // }
    fmt.Printf("unxi 时间戳=%v unxinano=%v\n",now.Unix(),now.UnixNano())
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run chapter05\demo18\mian.go
// unxi 时间戳=1773118145 unxinano=1773118145641368400