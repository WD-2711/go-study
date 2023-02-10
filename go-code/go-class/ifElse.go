package main


import "fmt"

func main(){
    num := 10
    // even 偶数
    if num % 2 == 0{
        fmt.Println("the number is even")
    } else {
        fmt.Println("odd")
    }
    
    //另一种写法
    if num1 := 10; num1 % 2 == 0{
        fmt.Println("the number is even")
    } else {
        fmt.Println("odd")
    }
    
    num2 := 99
    if num2 <= 50 {
        fmt.Println("number is less than 50")
    } else if num2 >= 51 && num <= 100 {
        fmt.Println("number is between 51 and 100")
    } else {
        fmt.Println("number is greater than 100")
    }
    
}