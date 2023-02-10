package main

import "fmt"
import "sync"

/* defer 应用的一个简单示例 */
func finished() {
    fmt.Println("Finished finding largest")
}
func largest(nums []int) {
    // 在largest函数返回前调用 defer 指定的 finished 函数
    defer finished()
    fmt.Println("Starting finding largest")
    max := nums[0]
    for _, v := range nums {
        if v > max {
            max = v
        }
    }
    fmt.Println("largest number in ", nums, "is", max)
}

/* 延迟方法 */
type person struct {
    firstName string
    lastName string
}
func (p person) fullName() {
    fmt.Printf("%s %s \n", p.firstName, p.lastName)
}

/* defer函数的实参取值 */
func printA(a int) {
    fmt.Println("value of a in deferred function", a)
}

/* defer 的实际应用 */
type rect struct {
    l int
    w int
}
func (r rect) area(wg *sync.WaitGroup) {
    // 重写后
    defer wg.Done()
    if r.l < 0 {
        fmt.Println("length wrong")
        // 重写前
        //wg.Done()
        return
    }
    if r.w < 0 {
        fmt.Println("width wrong")
        //wg.Done()
        return 
    }
    area := r.w * r.l
    fmt.Printf("rect %v area is %d\n", r, area)
    //wg.Done()
}

func main() {
    
    /* defer 应用的一个简单示例 */
//     nums := []int{78, 109, 2, 563, 300}
//     largest(nums)
    
    /* 延迟方法 */
//     p := person {
//         firstName: "clong",
//         lastName: "xie",
//     }
//     defer p.fullName()
//     fmt.Printf("welcome ")
    
    
    /* 
    defer的实参取值。
    defer不是在最后调用延迟函数的时候才取值，而是在defer处就完成了参数求值。
    */
//     a := 5
//     defer printA(a)
//     a = 10
//     fmt.Println("a's value", a)
    
    /*
    defer栈
    当有多个defer函数时，defer函数按照栈先进先出的顺序进行执行
    */
     
    /* defer 的实际应用*/
    // 用sync.WaitGroup做的，每次都写了一个wg.Done()，比较麻烦，用defer重写一下
    var wg sync.WaitGroup
    r1 := rect{-60, 80}
    r2 := rect{5, -67}
    r3 := rect{8, 9}
    rects := []rect{r1, r2, r3}
    for _, v := range rects {
        wg.Add(1)
        go v.area(&wg)
    }
    wg.Wait()
    fmt.Println("All go routines finished")
    
    
    
    
}