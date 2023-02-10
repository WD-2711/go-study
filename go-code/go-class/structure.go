
/*
命名的结构体
type Employee struct {
    first, second string
    age, salary int
}   

匿名的结构体，用一次
var Employee struct {
    first, second string
    age, salary int
}
*/

package main

import "fmt"


type Employee struct {
    f, l string
    age, salary int
}

// 匿名结构体
type Person struct {
    string
    int
}

// 嵌套结构体
type Addr struct {
    city, state string
}
type Human struct {
    name string
    age int
    addr Addr
}

// 嵌套结构体（提升字段）
type Addr2 struct {
    city, state string
}
type Human2 struct {
    name string
    age int
    Addr2
}

// 不可比较的结构体
type NotCompare struct {
    m map[string]int
}


func main(){
    emp1 := Employee{
        f: "clong",
        l: "xie",
        age: 20,
        salary: 66666,    
    }
    emp2 := Employee{"jinhe", "wu", 60, 80}
    fmt.Println(emp1, emp2)
    
    // 匿名结构体
    emp3 := struct {
        f, l string
        age, salary int
    }{
        f: "fei",
        l: "zhu",
        age: 66,
        salary:88888,
    }
    fmt.Println(emp3)
    // 访问结构体的字段
    fmt.Println(emp3.l)
    
    
    // 指向结构体的指针
    emp3_p := &emp3
    fmt.Println((*emp3_p).salary)
    // Go语言允许我们这么写，为了方便
    fmt.Println(emp3_p.salary)
    
    // 匿名结构体的使用
    p := Person{"xcl", 23}
    fmt.Println(p, p.string, p.int)
    
    // 嵌套结构体的使用
    var pp Human
    pp.name = "xcl"
    pp.age = 11
    pp.addr = Addr{
        city: "hefei",
        state: "Anhui",
    }
    fmt.Println(pp.name, pp.addr.city)
    
    
    // 当然，也可以这么写（提升字段）
    var pp2 Human2
    pp2.name = "xcl"
    pp2.age = 11
    pp2.Addr2 = Addr2{
        city: "hefei",
        state: "Anhui",
    }    
    fmt.Println(pp2.name, pp2.city)
    
    
    // 判断结构体的相等性
    p1 := Person{"xcl", 23}
    p2 := Person{"wjh", 20}
    if p1 == p2 {
        fmt.Println("equal")
    } else {
        fmt.Println("notEqual")
    }
    // 如果结构体含有不可比较的字段，则结构体不可比较
    /*
    nc1 := NotCompare{
        m: map[string]int{"xcl", 23} 
    }
    nc2 := NotCompare{
        m: map[string]int{"xc2", 23}
    } 
    if nc1 == nc2 {
        fmt.Println("equal")
    } else {
        fmt.Println("notEqual")
    }
    */
    
}