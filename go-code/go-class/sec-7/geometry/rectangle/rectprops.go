// rectprops.go
// 在rectangle文件夹中，所有文件都会以package rectangle作为开头

package rectangle

import "math"
import "fmt"

func init(){
    fmt.Println("rectangle package Initialized")
}

func Area(leng, wid float64) float64 {
    area := leng * wid
    return area
}

// Diagonal->对角线
func Diagonal(leng, wid float64) float64 {
    diagonal := math.Sqrt((leng * leng) + (wid * wid))
    return diagonal
}