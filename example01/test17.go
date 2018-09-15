package main
//(17)URL解析
//URL提供了一种统一访问资源的方式。
import (
	"net/url"
	"fmt"
	"strings"
)

func main() {
	// 将解析这个URL，它包含了模式，验证信息，主机，端口，路径，查询参数和查询片段
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// 解析URL，并保证没有错误
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// 可以直接访问解析后的模式
	fmt.Println(u.Scheme)

	// User包含了所有的验证信息，使用Username和Password来获取单独的信息
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// Host包含了主机名和端口，如果需要可以手动分解主机名和端口
	fmt.Println(u.Host)
	h := strings.Split(u.Host, ":")
	fmt.Println(h[0])
	fmt.Println(h[1])

	//这里我们解析出路径和`#`后面的片段
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// 为了得到`k=v`格式的查询参数，使用RawQuery。可以将查询参数解析到一个map里面。这个map为字符串作为key，
	// 字符串切片作为value。
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

}
