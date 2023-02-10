package main

import (
    "fmt"
    "math"
)

func main(){
    const a = 55
    // a = 99
    
    var b = math.Sqrt(4)
    // const b = math.Sqrt(4)
    fmt.Println(b)
    
    // Go是一个强类型的语言
    var defaultName = "sam"
    type myString string
    var newName myString = "alice"
    //newName = defaultName // 会报错，因为两个东西类型不一样
    fmt.Printf(defaultName, newName)
    
    // %T type %v value %d 整数
    var aa = 5.9/8
    fmt.Printf("\na's type %T value %v\n", aa, aa)
}