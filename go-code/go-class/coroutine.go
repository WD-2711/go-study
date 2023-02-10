package main

import "fmt"
import "time"

func hello() {
    fmt.Println("hello coroutine")
}

/* 同时启动多个 go 协程 */
func numbers() {
    for i := 1; i <= 5; i++ {
        time.Sleep(250 * time.Millisecond)
        fmt.Printf("%d ", i)
    } 
}
func alphabets() {
    for i := 'a'; i <= 'e'; i++ {
        time.Sleep(400 * time.Millisecond)
        fmt.Printf("%c ", i)
    }
}

func main(){
    
    /* 以下代码只会输出main function，因为hello协程来没来的及执行，程序就结束了 */
    // go hello()
    // fmt.Println("main function")
    
    
    /* 修复上述问题 */
    go hello()
    // time.Sleep(1 * time.Second)
    // fmt.Println("main function")
    
    
    /* 同时启动多个 go 协程 */
    go numbers()
    go alphabets()
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("main end")
}