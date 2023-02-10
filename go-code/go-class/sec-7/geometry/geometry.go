// geometry.go(几何学.go)
package main

import (
    "fmt"
    "geometry/rectangle"
    "log"
)

// 包级别的变量
var rectLen, rectWid float64 = 6, 7

// Init函数来检查长宽是否正确
func init(){
    println("main package Initialized")
    if rectLen < 0 {
        log.Fatal("length err")
    }
    if rectWid < 0 {
        log.Fatal("width err")
    }    
}


func main(){
    fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWid))
    fmt.Printf("diagonal of the rectangle %.2f\n", rectangle.Diagonal(rectLen, rectWid))
}