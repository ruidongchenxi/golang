package main

import (
	"fmt"
	"time"
	//"sync"
)
var number,letter= make(chan bool),make(chan bool)
//var wg sync.WaitGroup
func printNum(){
	i :=1
	for {
		<-number
		fmt.Printf("%d%d",i,i+1)
		i+=2
		letter<-true
	}
}
func printLetter(){
	i :=0
	str:="ABCDEFGHIJKLMNOPQSTUVWXYZ"
	for {
		<-letter
		if i+2>=len(str){
			fmt.Print(str[len(str)-1:])
			return
		}
		fmt.Print(str[i:i+2])
		i+=2
		number<-true
	}
}
func main(){

	go printNum()
	go printLetter()
	number <- true
	time.Sleep(time.Second*5)
}
