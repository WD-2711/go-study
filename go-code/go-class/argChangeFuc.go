package main

import (
    "fmt"
)

// 你看，这就是可变参数吧，...int意思就是int的数组
func find(num int, nums ...int) bool {
    found := false
    
    for _, v := range nums {
        if v == num{
            found = true
            return found
        }
    }
    return found
}

func changeStr(str ...string){
    str[0] = "cccc"
}


func main(){
    // 查找45是否在[89, 90, 67, 45, 109]
    res := find(45, 89, 90, 67, 45, 109)
    fmt.Println(res)

    
    // 查找45是否在[89, 90, 67, 109]
    res = find(45, 89, 90, 67, 109)
    fmt.Println(res)
    
    
    nums := []int{88, 99, 111}
    // 这是错误的，不能传切片，因为函数参数是 ...int
    // res = find(88, nums)
    // 但是这样是正确的，神奇不神奇？？
    res = find(88, nums...)
    fmt.Println(res)    
    
    // 神奇的事情，会输出[cccc bbbbbb]
    // 原因：当使用语法糖slice...后，slice会直接作为切片形式传入changeStr函数，此时str ...string就不是str的数组，而是切片
    slice := []string{"aaaaa", "bbbbbb"}
    changeStr(slice...)
    fmt.Println(slice)
}