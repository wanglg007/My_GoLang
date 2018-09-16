package main

/**
(3)基本类型
更明确的数字类型命名，⽀支持 Unicode，⽀支持常⽤用数据结构
 */

import (
	"math"
	"fmt"
)

func main() {
	a, b, c, d := 071, 0x1F, 1e9, math.MinInt16
	fmt.Println(a, b, c, d)
}
