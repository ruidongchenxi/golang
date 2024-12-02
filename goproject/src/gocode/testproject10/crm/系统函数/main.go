package main
import (
	"fmt"
    "strconv"
	"strings"
)	
func main(){
	str := "golang你好"//在golang中，汉字是utf-8字符集，一个汉字3个字节
	fmt.Println(len(str))
	//对字符串进行遍历：
	//方式1：利用键循环：for-range
	for i, value := range str {
		fmt.Printf("索引为：%d,具体的值为: %c \n",i,value)
	}
	//方式2
	r := []rune(str)
	for i :=0; i < len(r); i++{
		fmt.Printf("%c\n",r[i])
	}
	//字符串转整数
	num1,_ := strconv.Atoi("666")
	fmt.Println(num1)
	//整数转字符串
	str1 := strconv.Itoa(88)
	fmt.Println(str1)
	//查找子串是否在指定的字符串中，返回bool值
	str2 := strings.Contains("javaandgolang","go")
	fmt.Println(str2)
	//统计字符有几个指定的子串
	cout := strings.Count("golangandjava","ga")
	fmt.Println(cout)
	//不区分大小写字符串比较,返回bool值
	str3 := strings.EqualFold("go","GO")
	fmt.Println(str3)
	//区分大小写字符串比较,返回bool值
	fmt.Println("hello"=="Hello")
	//返回字串在字符串第一次出现的索引值，如果没有返回-1
    n1 := strings.Index("golangandjavaandpython","py")
	fmt.Println(n1)
	//字符串替换
	s1 := strings.Replace("goandjavagogo","go","golang",-1)
	fmt.Println(s1)
	//按照指定的某字符，为分割标识，将一个学符串拆分成字符串数组
	arr := strings.Split("go-python-java","-")
	fmt.Println(arr)
	fmt.Println(strings.ToLower("GO"))
	fmt.Println(strings.ToUpper("go"))
	fmt.Println(strings.TrimSpace("   go and java   "))
	fmt.Println(strings.Trim("~golang~","~"))
	fmt.Println(strings.TrimLeft("~golang~","~"))
	fmt.Println(strings.TrimRight("~golang~","~"))
	fmt.Println(strings.HasPrefix("http://java.sun.com/jsp/jstl/fmt","http"))
	fmt.Println(strings.HasSuffix("demo.png","png"))

}