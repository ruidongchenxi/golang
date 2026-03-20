package model
import "fmt"
type account struct {
	zh string
	ye float64
	ma string
}
func ChuShi(z string,y float64,m string ) *account {
	if len(z) < 6 || len(z) > 10 {
		return nil 
	}
	if y < 20{
		return nil 
	}
	if len(m) !=6 {
		return nil 
	}
	return &account{
		zh: z,
		ye: y,
		ma: m,
	}
}
// func ChuShi() *account{
// 	return &account{}
// }
func (A *account)SetZH(n string){
	if len(n)>=6&&len(n)<=10{
		A.zh=n
	} else {
		fmt.Println("账户长度不对")
		
	}
}
func (A *account)SetYe(n float64){
	if n>20{
		A.ye=n
	}else{
		fmt.Println("余额不正确")
	}
}
func (A *account)SetMa(n string){
	if len(n)==6{
		A.ma=n
 	}else{
		fmt.Println("密码位数不正确")
	}
}
func (A *account)GetZh() string{
	return A.zh
}
func (A *account)GetYe() float64{
	return A.ye
}
func (A *account)GetMa() string{
	return A.ma
}
