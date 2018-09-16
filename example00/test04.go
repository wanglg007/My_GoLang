package main

import "fmt"

/**
(5)字符串
字符串是不可变值类型，内部⽤用指针指向 UTF-8 字节数组。
•默认值是空字符串 ""。
•⽤用索引号访问某字节，如 s[i]。
•不能⽤用序号获取字节元素指针，&s[i] ⾮非法。
•不可变类型，⽆无法修改字节数组。
•字节数组尾部不包含 NULL。
 */
func main() {
	//使用索引号访问字符 (byte)
	s := "abc"
	println(s[0] == '\x61', s[1] == 'b', s[2] == 0x63)

	//使⽤用 "`" 定义不做转义处理的原始字符串，⽀支持跨⾏
	s1 := `a
b\r\n\x00
c`
	println(s1)

	//连接跨⾏行字符串时，"+" 必须在上⼀一⾏行末尾，否则导致编译错误
	s2 := "Hello, " +
		"World!"
	s3 := "Hello, "
	//+ "World!" // Error: invalid operation: + untyped string
	fmt.Println(s2, s3)

	//⽀支持⽤用两个索引号返回⼦子串。⼦子串依然指向原字节数组，仅修改了指针和⻓长度属性
	s00 := "Hello, World!"
	s10 := s[:5]  // Hello
	s20 := s[7:]  // World!
	s30 := s[1:5] // ello
	fmt.Println(s00, s10, s20, s30)

	//要修改字符串，可先将其转换成 []rune 或 []byte，完成后再转换为 string。无论哪种转换，都会重新分配内存，并复制字节数组。
	ss := "abcd"
	bs := []byte(ss)
	bs[1] = 'B'
	println(string(bs))
	u := "电脑"

	us := []rune(u)
	us[1] = '话'
	println(string(us))

	//用 for 循环遍历字符串时，也有 byte 和 rune 两种⽅方式
	sm := "abc汉字"
	for i := 0; i < len(sm); i++ { // byte
		fmt.Printf("%c,", s[i])
	}
	fmt.Println()
	for _, r := range s { // rune
		fmt.Printf("%c,", r)
	}
}
