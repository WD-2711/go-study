package main

import "fmt"
import "sync"
import "time"

var x = 0
func increment(wg *sync.WaitGroup) {
    time.Sleep(1000 * time.Millisecond)
    x = x + 3
    wg.Done()
}

/* 使用Mutex修复竞争态条件的问题 */
func increment_2(wg *sync.WaitGroup, m *sync.Mutex){
    m.Lock()
    x = x + 1
    m.Unlock()
    wg.Done()
}

/* 使用信道修复竞争态条件的问题 */
func increment_3(wg *sync.WaitGroup, ch chan bool) {
    ch <- true
    x = x + 1
    <- ch
    wg.Done()
}



func main(){
    /* 竞争条件的例子，不知道为啥一直复现不出来，感觉应该是Go版本太新了*/
    //var w sync.WaitGroup
    //for i := 0; i < 1000; i++ {
    //    w.Add(1)
    //    go increment(&w)
    //}
    //w.Wait()
    //fmt.Println("final x", x)
    
    /* 使用Mutex修复竞争态条件的问题 */
    //var w sync.WaitGroup
    //var m sync.Mutex
    //for i := 0; i < 1000; i++ {
    //    w.Add(1)
    //    go increment_2(&w, &m) // 要传递Mutex的地址，要是不传地址，那每个协程都会得到mutex的拷贝，竞争态条件还是会发生
    //}
    //w.Wait()
    //fmt.Println("final x", x)
    
    /* 还可以用信道来解决竞争态问题 */
    var w sync.WaitGroup
    ch := make(chan bool, 1)
    for i := 0; i < 1000; i++ {
        w.Add(1)
        go increment_3(&w, ch) // 由于通道容量是1，因此满了之后会阻塞，因而解决了竞争态条件的问题
    }
    w.Wait()
    fmt.Println("final ", x)
}