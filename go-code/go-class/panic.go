package main

import "fmt"
import  "runtime/debug"

// func name(first *string) { 
//     defer fmt.Println("deferred call in name")
//     if first == nil {
//         panic("first can't nil")
//     }
//     fmt.Printf("%s", first)
//     return
// }

/* recover函数使用 */
// func recoverName() {
//     if r := recover(); r != nil {
//         fmt.Println("recovered from ", r)
//     }
// }
// func name(f *string) {
//     // 在此用了recover函数
//     defer recoverName()
//     if f == nil {
//         // 产生了panic的错误
//         panic("runtime error")
//     }
//     fmt.Printf("%s %s\n", *f)
// } 


/* 运行时panic */
// func a() {
//     n := []int{5, 7, 4}
//     fmt.Println(n[3]) // 这会出现数组溢出
//     fmt.Println("normal return from a")
// }

/* 恢复一个运行时panic */
func r() {
    if r:= recover(); r != nil {
        fmt.Println("Recovered", r)
        debug.PrintStack()
    }
}
func a() {
    defer r()
    n := []int{5, 7, 4}
    fmt.Println(n[3])
    fmt.Println("normally returned from a")
}

func main() {
    
    /*
    下面的示例讲了defer函数与panic运行的先后顺序
    1. 运行name函数的defer
    2. 运行了main函数中的defer
    3. 运行panic函数
    */
//     defer fmt.Println("deferred call in main")
//     name(nil)
//     fmt.Println("main end")
    
    /*
    recover用于重新获得panic协程的控制
    只有在延迟函数内部调用recover才有用
    然而，recover只能用于恢复此协程的panic，不能恢复对于不同协程的panic
    */
//     defer fmt.Println("deferred call in main")
//     name(nil)
//     fmt.Println("main end")
    
    /*
    运行时panic。
    啥是运行时panic，例如数组越界，相当于调用了go内置函数的panic。
    */
//     a()
//     fmt.Println("normally return from main")
    // 上述代码会出现panic
    
    /* 恢复一个运行时panic */
    a()
    fmt.Println("normally main end")
    
    /*
    但是上述恢复运行时panic会丢失对堆栈的追踪（就是详细的错误信息），那么我们怎么打印出堆栈追踪呢
    在r() 中调用debug.PrintStack()即可
    */
    
    
    
}