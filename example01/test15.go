package main
//(15)String与Byte切片之间的转换
// String转换到Byte数组时，每个byte(byte类型其实就是uint8)保存字符串对应字节的数值。注意Go的字符串是UTF-8编码的，每个字符
// 长度是不确定的，一些字符可能是1、2、3或者4个字节结尾
import "fmt"

func main() {
	s1 := "abcd"
	b1 := []byte(s1)
	fmt.Println(b1)

	s2 := "中文"
	b2 := []byte(s2)
	fmt.Println(b2)

	r1 := []rune(s1)
	fmt.Println(r1)

	r2 := []rune(s2)
	fmt.Println(r2)

}
