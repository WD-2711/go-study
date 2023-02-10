package main

import "fmt"
import "time"

/* 信道是阻塞的，编写如下代码使得Go主协程等待hello协程结束 */
func hello(done chan bool) {
    fmt.Println("hello coroutine")
    // true 和 false 都ok
    time.Sleep(4 * time.Second)
    done <- false
}

/* 信道的另一个例子 */
// 例如有123，那么计算 1*1+2*2+3*3
func calcSquares(number int, op chan int) {
    sum := 0
    for number != 0 {
        digit := number % 10
        sum += digit * digit
        number /= 10
    }
    op <- sum
}
// 例如有123，那么计算 1*1*1+2*2*2+3*3*3
func calcCubes(number int, op chan int) {
    sum := 0
    for number != 0 {
        digit := number % 10
        sum += digit * digit * digit
        number /= 10
    }
    op <- sum    
}

/* 单向信道 */
func sendData(ch chan<- int) {
    ch <- 10
}

/* 关闭信道 & for age 遍历信道 */
func producer(ch chan int) {
    for i := 0; i < 10; i++ {
        ch <- i
    }
    close(ch)
}

/* 优化上述calcSquares与calcCubes的代码，代码复用 */
func digits(number int, ch chan int) {
    for number != 0 {
        digit := number % 10
        ch <- digit
        number /= 10
    }
    close(ch)
}
func calcSquares_2(number int, ch chan int) {
    sum := 0
    ch2 := make(chan int)
    go digits(number, ch2)
    for d := range ch2 {
        sum += d * d
    }
    ch <- sum
}
func calcCubes_2(number int, ch chan int) {
    sum := 0
    ch2 := make(chan int)
    go digits(number, ch2)
    for d := range ch2 {
        sum += d * d * d
    }
    ch <- sum
}

func main() {
    /* 简单的信道 */
    var a chan int // int类型的信道
    if a == nil {
        fmt.Println("channel a is nil")
        fmt.Printf("Type of a is %T\n", a) // 在 make 前后都是 chan int 类型
        a = make(chan int)
        fmt.Printf("Type of a is %T\n", a)
    }
    
    /* 信道的发送与接收 */
    // b = make(chan int)
    // data := <- b // 读取信道b
    // b <- data // 写入信道b
    
    
    /* 信道是阻塞的，编写如下代码使得Go主协程等待hello协程结束 */
    // done := make(chan bool)
    // go hello(done)
    // <- done
    // fmt.Println("main end")
    
    /* 信道的另一个例子 */
    number := 589
    sq := make(chan int)
    cu := make(chan int)
    go calcSquares(number, sq)
    go calcCubes(number, cu)
    squares, cubes := <- sq, <- cu
    fmt.Println("Final ", squares + cubes)
    
    /* 死锁 */
    // 光写数据了，没读数据
    // ch := make(chan int)
    // ch <- 5
    
    /* 单向信道 */
    // 只是往信道里读，所以fmt.Println(<-ch)报错，但是这样的单项信道有啥意义呢？
    ch := make(chan<- int)
    go sendData(ch)
    // fmt.Println(<-ch)
    
    /* 信道转换，只能把双向信道转成单向信道，但是从单向信道到双向信道是不行滴 */
    ch1 := make(chan int)
    go sendData(ch1) // 双向信道转换成单向信道
    fmt.Println(<-ch1)
    
    /* 关闭信道 & for age 遍历信道 */
    ch2 := make(chan int)
    go producer(ch2)
    for {
        v, ok := <- ch2
        if ok == false {
            break
        }
        fmt.Println(v, ok)
    }
    // 也可以这么写
    ch3 := make(chan int)
    go producer(ch3)
    for v := range ch3 {
        fmt.Println(v)
    }
    
    /* 优化上述calcSquares与calcCubes的代码，代码复用 */
    number_2 := 589
    ch4 := make(chan int)
    go calcSquares_2(number_2, ch4)
    ch5 := make(chan int)
    go calcCubes_2(number_2, ch5)
    fmt.Println(<-ch4 + <-ch5)
}