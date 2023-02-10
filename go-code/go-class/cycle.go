package main

import "fmt"

func main(){
    for i := 1; i <= 10; i++ {
        fmt.Printf(" %d", i)
    }
    fmt.Printf("\n")
    
    // break语句, continue类似
    for i  := 1; i <= 10; i++ {
        if i > 5 {
            break
        }
        fmt.Printf("%d ", i)
    }
    fmt.Printf("\nloop end\n")
    
    // 另一种写法
    i := 0
    for i <= 10 {
        fmt.Printf("%d ", i)
        i += 2
    }
    fmt.Printf("\n")
    
    for i := 1; i <= 10; i++ {
        fmt.Printf("%d * %d = %d\n", i + 9, i, i * (i + 9))
    }
    
}