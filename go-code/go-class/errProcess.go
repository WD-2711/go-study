package main

import "fmt"
// import "os"
// import "net"
import "path/filepath"


/* 
错误类型的表示 
error是一个接口
*/


/* 直接比较 */

func main() {
    /* 错误类型例子 */
//     f, err := os.Open("/test.txt")
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//     fmt.Println(f.Name(), "open good")
    // 由于没有这个文件，所以会报错： open /test.txt: No such file or directory
    
    /* 更详细的看看这个错误 */
//     f, err := os.Open("/test.txt")
//     // 表示确实是 PathError 这个错误
//     if err, ok := err.(*os.PathError); ok {
//         fmt.Println(err.Path, "failed to open")
//         return
//     }
//     fmt.Println(f.Name(), "open good")
    
    /* 错误类型的表示，以 DNSError 为例 */
    /*
    表示这个错误是怎么产生的，是由Timeout产生的，还是Temporary产生的
    type DNSError struct {
        ...
    }
    func (e *DNSError) Error() string {
        ...
    }
    func (e *DNSError) Timeout() string {
        ...
    }
    func (e *DNSError) Temporary() string {
        ...
    }
    */
    /* 下面写一个DNSerror的例子*/
    // 该错误既不是临时性错误，也不是超时引发的，
    /*
    addr, err := net.LookupHost("golangbot123.com")
    if err, ok := err.(*net.DNSError); ok {
        if err.Timeout() {
            fmt.Println("operation timed out")
        } else if err.Temporary() {
            fmt.Println("temporary error")
        } else {
            fmt.Println("error:", err)
        }
        return
    }
    fmt.Println(addr)
    */
   
    /* 直接比较：这是获取错误信息的另一种方式，是直接与error类型的变量作比较 */
    /*
    filepath.Glob("*.txt")
    查找当前目录下所有具有.txt扩展名的所有文件并打印他们的名字
    */
    // 查找当前目录下含有'['的文件
    files, _ := filepath.Glob("[")
//     if error != nil && error == filepath.ErrBadPattern {
//         fmt.Println(error)
//         return
//     }
    fmt.Println("matched files", files)
    /*
    会报syntax error in pattern错误
    因为ErrBadPattern的定义如下：
    var ErrBadPattern = errors.New("syntax error in pattern")
    */
    
    /* 
    不可忽略错误。
    绝不要忽略错误，忽视错误会带来问题。
    如果不加上述错误判断，直接判断的话会直接输出[]，代表没有匹配，但是也不直接报错。
    */
    
    
    
}