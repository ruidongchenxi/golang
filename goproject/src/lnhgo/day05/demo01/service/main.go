package main

import (
	"encoding/json"
	"fmt"
	//"html/template"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		_=r.ParseForm()//解析参数
		fmt.Println("path:", r.URL.Path)
		// fmt.Println("scheme:", r.URL.Scheme)
		// fmt.Println(r.Form)
		a,_:=strconv.Atoi(r.Form["a"][0])
		b,_:=strconv.Atoi(r.Form["b"][0])
		w.Header().Set("Content-Type","application/json")
		jData,_:=json.Marshal(map[string]int{
			"data":a+b,
		})
		_,_=w.Write(jData)
	})
	_=http.ListenAndServe(":8000",nil)
}