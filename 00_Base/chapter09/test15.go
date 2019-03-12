package main

//(9.15)接口案例
import "fmt"

type Usb15 interface {
	Say()
}
type Stu15 struct {
}

func (this *Stu15) Say() {
	fmt.Println("Say()")
}
func main() {
	var stu Stu15 = Stu15{}
	// 错误！ 会报 Stu15类型没有实现Usb15接口 ,
	// 如果希望通过编译,  var u Usb15 = &stu
	//var u Usb15 = stu
	var u Usb15 = &stu
	u.Say()
	fmt.Println("here", u)
}
