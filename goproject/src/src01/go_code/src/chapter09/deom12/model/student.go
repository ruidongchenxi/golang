package model
//定义结构体
type student struct{
	Name string
	score float64
}
//因为student结构体首字母是小写的，因此只能在model包使用
//通过工厂模式解决,接收两个参数返回一个结构体地址值
func NewStudent(n string,s float64) *student{
	return &student{
		Name: n,
		score: s,
	}
}
func Newget(n *student) float64{ 
	return n.score
}
func (t *student)Test() float64{
	return t.score
}