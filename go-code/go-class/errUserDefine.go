/*
error包中New函数实现：
package errors
func New(text string) error {
    return &errorString{text}
}
type errorString struct {
    s string
}
func (e *errorString) Error() string {
    return e.s
}
*/

package main

import (
    _ "errors"
    "fmt"
    _ "math"
)

/* 1. 新建 Error 的一个简单示例 */
// func circleArea(r float64) (float64, error) {
//     if r < 0 {
// //         return 0, errors.New("Area could't less than 0")
//         /* 2. 使用Errorf给错误添加更多信息 */
//         return 0, fmt.Errorf("Area calculation failed, %0.2f is less than zero", r)
//     }
//     return math.Pi * r * r, nil
// }

/* 3. 使用结构体类型和字段提供错误的更多信息 */
// type areaError struct {
//     err string // 存储与错误有关的信息
//     r float64 // 存储与错误有关的半径
// }
// func (e *areaError) Error() string {
//     return fmt.Sprintf("radius %0.2f: %s", e.r, e.err)
// }
// func circleArea(r float64) (float64, error) {
//     if r < 0 {
//         return 0, &areaError{"radius is negative", r}
//     }
//     return math.Pi * r * r, nil
// }

/* 4. 使用结构体类型的方法来输出错误的练习 */
// type rectangle struct {
//     l int
//     w int
// }
// type areaError struct {
//     err string
//     l int
//     w int
// }
// func (e *areaError) Error() string {
//     return fmt.Sprintf("rectangle width %d and length %d: %s", e.w, e.l, e.err)
// }
// func (r *rectangle) calArea() (int, error){
//     if r.l < 0 {
//         return 0, &areaError{"length less 0", r.l, r.w}
//     }
//     if r.w < 0 {
//         return 0, &areaError{"width less 0", r.l, r.w}   
//     }
//     return r.l * r.w, nil
// }

/* 5. 4的标准示例代码 */
type areaError struct {
    err string
    l float64
    w float64
}
// 如果真要成为一个error，那么一定要实现error方法
func (e *areaError) Error() string {
    return e.err
}
func (e *areaError) lbad() bool {
    return e.l < 0
}
func (e *areaError) wbad() bool {
    return e.w < 0
}
func rectArea(l, w float64) (float64, error) {
    err := ""
    if l < 0 {
        err += "l is less than 0"
    }
    if w < 0 {
        if err == "" {
            err = "w is less than 0"
        } else {
            err += ", w is less than 0"
        }
    }
    if err != "" {
        return 0, &areaError{err, l, w}
    }
    return l * w, nil
}

func main() {
    /* 1. 新建 Error 的一个简单示例 */
    /* 2. 使用Errorf给错误添加更多信息 */
    /* 3. 还可以使用结构体类型和字段提供错误的更多信息 */
    /* 2中处理的不好的地方：如果我们想访问引发错误的半径，则必须访问错误字符串，不太好，所以用3中的结构体处理一波 */
//     radius := -20.0
//     area, err := circleArea(radius)
//     if err != nil {
//         if err, ok := err.(*areaError); ok {
//             fmt.Println(err.r, err) // 如果是err就运行areaError的Error方法
//             fmt.Println(err.r, err.err) // 如果是err.err的话就输出areaError的err字符串
//         }
//         return
//     }
//     fmt.Printf("Area of circle %0.2f\n", area)
    
    /* 4. 使用结构体类型的方法来输出错误的练习 */
//     rec := rectangle{20, -20}
//     area, err := rec.calArea()
//     if err != nil {
//         if err, ok := err.(*areaError); ok {
//             fmt.Println(err.l, err.w, err)
//         }
//         return
//     }
//     fmt.Println(area)
    
    /* 5. 4的标准示例代码 */
    l, w := -5.0, -9.0
    _, err := rectArea(l, w)
    if err != nil {
        if err, ok := err.(*areaError); ok {
            if err.lbad() {
                fmt.Println("11")
            }
            if err.wbad() {
                fmt.Println("22")
            }
            return
        }
        fmt.Println(err)
        return
        
    }
    
    
}