package main
//(5)Go 自定义排序
//自定义排序 - 希望按照字符串的长度来对一个字符串数组排序而不是按照字母顺序来排序。
import (
	"sort"
	"fmt"
)

// 为了能够使用自定义函数来排序，需要一个对应的排序类型，比如这里为内置的字符串数组定义了一个别名ByLength
type ByLength []string

// 实现了sort接口的Len，Less和Swap方法。这样就可以使用sort包的通用方法Sort。Len和Swap方法的实现在不同的类型之间大致
// 都是相同的，只有Less方法包含了自定义的排序逻辑，这里希望以字符串长度升序排序
func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i,j int) {
	s[i],s[j] = s[j],s[i]
}

func (s ByLength) Less(i,j int) bool {
	return len(s[i]) < len(s[j])
}

// 把需要进行自定义排序的字符串类型fruits转换为ByLength类型，然后使用sort包的Sort方法来排序
func main()  {
	fruits := []string{"peach","banana","kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}


