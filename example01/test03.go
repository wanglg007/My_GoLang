package main
//(3)Go 字符串操作函数
//strings 标准库提供了很多字符串操作相关的函数。
import (
	"fmt"
	"strings"
)

//strings 标准库提供了很多字符串操作相关的函数。

//给fmt.Println起个别名，因为下面会多处使用
var p = fmt.Println

func main()  {
	// 下面是strings包里面提供的一些函数实例。注意这里的函数并不是string对象所拥有的方法，这就是说使用这些字符串操作函数的时候
	// 你必须将字符串对象作为第一个参数传递进去。
	p("Contains:",strings.Contains("test","es"))
	p("Count:",strings.Count("test","t"))
	p("HasPrefix:",strings.HasPrefix("test","te"))
	p("HasSuffix:",strings.HasSuffix("test","st"))
	p("Index:",strings.Index("test","e"))
	p("Join:",strings.Join([]string{"a","b"},"-"))

	p("Repeat:",strings.Repeat("a",5))
	p("Replace:",strings.Replace("foo","o","0",-1))
	p("Replace:",strings.Replace("foo","o","0",1))
	p("Split:",strings.Split("a-b-c-d-e","-"))

	p("ToLower:",strings.ToLower("TEST"))
	p("ToUpper:",strings.ToUpper("test"))
	p()

	// 这里还有两个字符串操作方法，一个是获取字符串长度，另外一个是从字符串中获取指定索引的字符
	p("Len:",len("hello"))
	p("Char:","hello"[1])

}
