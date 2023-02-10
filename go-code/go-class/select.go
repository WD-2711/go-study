package main

import "fmt"
import "time"

/* 一个select小例子 */
func server1(ch chan string) {
    time.Sleep(6 * time.Second)
    ch <- "from server1"
}
func server2(ch chan string) {
    time.Sleep(3 * time.Second)
    ch <- "from server2"
}

/* 当case没有就绪时，可以执行默认情况Default Case，来防止select语句一直阻塞 */
func process(ch chan string) {
    time.Sleep(10500 * time.Millisecond)
    ch <- "process successful"
}

/* 当多个 case 准备就绪时，select会随机选取其中一个来执行 */
func server3(ch chan string) {
    ch <- "from server3"
}
func server4(ch chan string) {
    ch <- "from server4"
}

func main() {
      
    /* 一个select小例子 */
    //op1 := make(chan string)
    //op2 := make(chan string)
    //go server1(op1)
    //go server2(op2)
    //select {
    //    case s1 := <-op1:
    //        fmt.Println(s1)
    //    case s2 := <-op2:
    //        fmt.Println(s2)
    //}
    
    
    /* 当case没有就绪时，可以执行默认情况Default Case，来防止select语句一直阻塞 */
    //ch := make(chan string)
    //go process(ch)
    //for {
    //    time.Sleep(1000 * time.Millisecond)
    //    select {
    //        case v := <-ch:
    //            fmt.Println("received value ", v)
    //            return // 跳出for循环, break办不了
    //        default:
    //            fmt.Println("no received")
    //    }
    //}
    
    /* 当没有输出到信道时，如果没有default情况，会报panic错误 */
    
    /* 当多个 case 准备就绪时，select会随机选取其中一个来执行 */
    op3 := make(chan string)
    op4 := make(chan string)
    go server3(op3)
    go server4(op4)
    time.Sleep(1 * time.Second)
    select {
        case s1 := <-op3:
            fmt.Println(s1)
        case s2 := <-op4:
            fmt.Println(s2)
    }
    
    /*
    空的select会导致死锁。
    select {} 因为：除非有case执行，否则select语句就会一直阻塞。
    */

    
    
}