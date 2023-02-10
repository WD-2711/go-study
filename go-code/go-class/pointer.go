package main

import (
    "fmt"
)

func change(bb *int){
    *bb = 60
}

func modify(a2 []int){
    a2[0] = 1
}

func main(){
    b := 255
    var a *int = &b
    fmt.Printf("%T\n", a) // a的类型，a指向b
    fmt.Println(a) // b的地址
    
    
    aa := 25
    var bb *int
    if bb == nil {
        fmt.Println(bb)
        bb = &aa
        fmt.Println(bb)
        fmt.Println(*bb) // 指针的解引用
    }
    
    a1 := 58
    b1 := &a1
    change(b1)
    fmt.Println(a1)
    
    // 可以看到，指针应该可以更改数组的值，但是更推荐用切片
    a2 := [3]int{8, 9, 10}
    modify(a2[:])
    fmt.Println(a2)
    
    // 与C不同的是，go不支持指针运算
    a3 := 1
    b3 := &a3
    // 会报错
    // b3 += 1
}