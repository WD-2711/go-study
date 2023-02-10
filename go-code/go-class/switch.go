package main

import "fmt"

func main(){
    finger := 4
    
    switch finger {
        case 1: 
            fmt.Println("One")
        case 2:
            fmt.Println("Two")
        default:
            fmt.Println("else")
        
    }
    // 另一种写法
    switch f := 8; f {
        case 1: 
            fmt.Println("One")
        case 2:
            fmt.Println("Two")
        default:
            fmt.Println("else")
    }
    
    // 多表达式判断，表达式指的是"a", "e", "i", "o", "u"
    switch le := "i"; le {
        case "a", "e", "i", "o", "u":
            fmt.Println("vowel") // vowel元音
        default:
            fmt.Println("else")
    }
    
    // 无表达式判断
    num := 75
    switch {
        case num >= 0 && num <= 50:
            fmt.Println("0-50")
        case num >= 51 && num <= 100:
            fmt.Println("51-100")
        case num >= 101:
            fmt.Println("100-")
    }
    
    switch num1 := 30; {
        case num1 < 50:
            fmt.Println("<50")
            fallthrough
        case num1 < 70:
            fmt.Println("<70") 
    }
    
    
}