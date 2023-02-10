package main

import "fmt"


func changeMap(mm map[string]int){
    _, ok := mm["aaa"]
    if ok == true{
        mm["aaa"] = 888
    }
}

func main(){

    var m map[string]int 
    if m == nil {
        // 对map进行初始化
        m = make(map[string]int)
    }
    
    m1 := make(map[string]int)
    m1["a"] = 1
    m1["b"] = 2
    fmt.Println(m1)
    
    // 对map进行初始化值
    m2 := map[string]int{
        "a": 1,
        "b": 2, // 这里也要加","，但是不知道为什么
    }
    fmt.Println(m2)
    // 如果取一个不存在的元素会返回0
    fmt.Println(m2["c"])
    // 判断m2中是否有"d"索引, 如果ok = false则说明没有，记住go是强类型，所以ok==1会报错
    _, ok := m2["d"]
    if ok == true{
        fmt.Println("yes")
    } else {
        fmt.Println("no")
    }
    // m2的元素遍历
    for k, v := range m2{
        fmt.Printf("%s %d\n", k, v)
    }
    // 删除m2中的某个元素，map的长度为len(m2)
    delete(m2, "a")
    fmt.Println(m2)
    m2["a"] = 1
    
    
    
    // map与slice一样都是引用类型，他们都指向底层的数据结构
    m3 := map[string]int{
        "aaa": 122121,
        "bvbbb": 6666,
    }
    fmt.Println(m3)
    changeMap(m3)
    fmt.Println(m3)
    
    
    // map的相等性判断，==只能用来检查map是否为nil，不能判断两个map是否相等
    
}