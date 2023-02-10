package main

import (
    "fmt"
    "structs/computer"
)

func main(){
    var spec computer.Spec
    spec.Maker = "apple"
    spec.Price = 50000
    // spec.model = "Mac mini"
    // 访问不到哦
    fmt.Println(spec)
}