package main

import "fmt"
import "time"
import "sync"
import "math/rand"

/* 缓冲信道的另一个示例 */
func write(ch chan int) {
    for i := 0; i < 5; i++ {
        ch <- i
        fmt.Printf("write %d to ch\n", i)
    }
    close(ch)
}

/* WaitGroup示例 */
func process(i int, wg *sync.WaitGroup) {
    fmt.Println("start routine", i)
    time.Sleep(2 * time.Second)
    fmt.Println("end routine", i)
    wg.Done()
}

/* 工作池举例 （重点）*/
// 要完成的工作Job
type Job struct {
    id int
    randomno int
}
// 工作结果放入到Result中
type Result struct {
    job Job
    sumofdigits int
}
var jobs = make(chan Job, 10)
var results = make(chan Result, 10)
// 计算整数的每一位之和
func digits(number int) int {
    sum := 0
    no := number
    for no != 0 {
        digit := no % 10
        sum += digit
        no /= 10
    }
    time.Sleep(2 * time.Second)
    return sum
}
// 创建工作协程
func worker(wg *sync.WaitGroup) {
    for job := range jobs {
        output := Result{job, digits(job.randomno)}
        results <- output
    }
    wg.Done()
}
// 协程的工作池
func createWorkerPool(noOfWorkers int) {
    var wg sync.WaitGroup
    for i := 0; i< noOfWorkers; i++ {
        wg.Add(1)
        go worker(&wg)
    }
    wg.Wait()
    close(results)
}
// 生成好多jobs
func allocate(noOfJobs int) {
    for i := 0; i < noOfJobs; i++ {
        randomno := rand.Intn(999)
        job := Job{i, randomno}
        jobs <- job
    }
    close(jobs)
}
// 打印输出的result
func result(done chan bool) {
    for result := range results {
        fmt.Printf("job id %d, input random no %d, sum of digits %d\n",result.job.id, result.job.randomno, result.sumofdigits)
    }
    done <- true
}

func main() {
    ch := make(chan string, 2) // 缓冲信道，容量为2
    ch <- "xcl"
    ch <- "wjh"
    fmt.Println(<- ch)
    fmt.Println(<- ch)
    
    /* 缓冲信道的另一个示例 */
    // ch1 := make(chan int, 2)
    // go write(ch1)
    // time.Sleep(2 * time.Second)
    // for v := range ch1 {
    //     fmt.Printf("read %d from channel\n", v)
    //     time.Sleep(2 * time.Second)
    // }
    
    /* WaitGroup示例 */
    // no := 3
    // var wg sync.WaitGroup
    // for i := 0; i < no; i++ {
    //     wg.Add(1)
    //     go process(i, &wg) // 开了协程，所以同时运行
    // }
    // wg.Wait()
    // fmt.Println("All routines end")
    
    /* 工作池举例 （重点）*/
    st := time.Now()
    noOfJobs := 100
    go allocate(noOfJobs)
    done := make(chan bool)
    // 遍历的时候，for result := range results，应该是被阻塞
    go result(done)
    noOfWorkers := 20
    createWorkerPool(noOfWorkers)
    <- done
    et := time.Now()
    fmt.Println("total time ", et.Sub(st).Seconds())
    
    
}