package main

//(9.2)封装案例
import (
	"fmt"
	"github.com/wanglg007/My_GoLang/00_Base/chapter09/model"
)

func main() {
	//创建一个account变量
	account := model.NewAccount("jzh11111", "000", 40)
	if account != nil {
		fmt.Println("创建成功=", account)
	} else {
		fmt.Println("创建失败")
	}
}