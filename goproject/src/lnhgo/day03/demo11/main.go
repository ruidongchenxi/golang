package main

import "fmt"

type MyWriter interface {
	Reader(string)
}
type MyReader interface {
	Write(string)
}
type MyReadWriter interface {
	MyReader
	MyWriter
	ReadWrite()
}
type SreadWriter struct{}

func (s *SreadWriter)Write(string) {
	fmt.Println("implement me")
}
func (s *SreadWriter)Reader(string){
	fmt.Println("implement Reader")
}
func (s *SreadWriter)ReadWrite(){
	fmt.Println("ReadWrite()")
}
func main() { 
	var mrw MyReadWriter = &SreadWriter{}
	mrw.Reader("t")
 }