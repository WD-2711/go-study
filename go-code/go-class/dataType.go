package main

import (
    "fmt"
    "unsafe"
)

func main() {
    var a int = 89
    b := 95
    fmt.Println("value of a is", a, "and b is", b)
    fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a)) // a 的 type，由于a是int且为64位系统，所以占8个字节。
    fmt.Printf("\ntype of b is %T, size of b is %d\n", b, unsafe.Sizeof(b))
    
    c, d := 5.67, 8.97
    fmt.Println("sum", c + d, "diff", c - d)
    
    // 复数
    c1 := complex(5, 7)
    c2 := 8+ 27i
    cadd := c1 + c2
    fmt.Println("sum: ", cadd)
    cmul := c1 * c2
    fmt.Println("product:", cmul)
    
    // String类型
    first := "head"
    last := "tail"
    name := first + " " + last
    fmt.Println("My name is", name)
    
    // 类型转换
    i := 10
    var j float64 = float64(i)
    // var j float64 = i
    fmt.Println("j", j)
}
