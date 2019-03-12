package main

//(9.2)封装
import (
	"fmt"
	"github.com/wanglg007/My_GoLang/00_Base/chapter09/model"
)

func main() {

	p := model.NewPerson("smith")
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(p)
	fmt.Println(p.Name, " age =", p.GetAge(), " sal = ", p.GetSal())

}
