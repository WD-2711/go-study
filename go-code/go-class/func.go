package main

import "fmt"


func calculateBill(price, no int) int {
    var totalPrice = price * no
    return totalPrice
}

/*
func rectProps(length, width float64)(float64, float64){
    var area = length * width
    var perimeter = (length + width) * 2
    return area, perimeter
}
*/
// 也可以这么写
func rectProps(length, width float64)(area float64, perimeter float64){
    area = length * width
    perimeter = (length + width) * 2
    return
}


func main(){
    
    // 单返回值
    price, no := 90, 6
    totalPrice := calculateBill(price, no)
    fmt.Println("Total price is", totalPrice)
    
    // 多返回值, %f float64
    area, perimeter := rectProps(10.8, 5.6)
    fmt.Printf("Area %f Perimeter %f\n", area, perimeter)
    
    // 空白符
    newArea, _ := rectProps(10.8, 5.6)
    fmt.Printf("Area %f\n", newArea)
}