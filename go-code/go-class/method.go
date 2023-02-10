
package main

import (
    "fmt"
    "math"
)

type Employee struct {
    name string
    salary int
    currency string
}

// displaySalary() 方法讲 Employee 作为接收器类型
func (e Employee) displaySalary() {
    fmt.Printf("%s %d %s\n",e.name, e.salary, e.currency)
}

type Rectangle struct {
    l int
    w int
}

type Circle struct {
    r float64
}

// 定义在Rectangle上的方法
func (r Rectangle) Area() int {
    return r.l * r.w
}

// 定义在Circle上的方法
func (c Circle) Area() float64 {
    return math.Pi * c.r * c.r
}

// area函数，参数为值
func area(r Rectangle) int {
    return r.l * r.w
}

type Employee1 struct {
    name string
    age int
}

// 值接收器
func (e Employee1) changeName(newName string){
    e.name = newName
}

// 指针接收器
func (e *Employee1) changeAge(newAge int){
    e.age = newAge
}

// 非结构体上定义的方法 
// 此时要为类型创建一个类型别名，在使用接收器
type MyInt int
func (a MyInt)MyMethod(b MyInt) MyInt {
    return a+b
}

func main() {
    
    emp1 := Employee {
        name: "xcl",
        salary: 6666666,
        currency: "$",
    }
    // 这里是不是就能看出来函数与方法的区别啦
    emp1.displaySalary()
    
    // 定义在不同对象上的方法
    r := Rectangle{
        l: 50,
        w: 60,
    }
    c := Circle{
        r: 3.33,
    }
    
    fmt.Println(r.Area())
    
    fmt.Println(c.Area())
    
    // 方法的值接收器与指针接收器。（上面的都是值接收器）
    // 可以看到，此时传值的话是不能改变结构体内部的值，只有传指针才行
    e := Employee1{
        name: "xcl",
        age: 23,
    }
    fmt.Println(e)
    e.changeName("wjh")
    fmt.Println(e)
    (&e).changeAge(3333)
    fmt.Println(e)
    // 当然咯，也可以这么写(把&去掉)
    e.changeAge(444)
    fmt.Println(e)    
    
    // 方法中的值接收器 与 函数中的值参数
    rr1 := Rectangle{
        l: 6,
        w: 7,
    }
    fmt.Println("value args for value function",area(rr1))
    // fmt.Println(area(&rr1)) // 会报错
    fmt.Println("value args for value method",rr1.Area())
    fmt.Println("pointer args for value method",(&rr1).Area()) // go语言为了方便
    
    
    // 上面全是在结构体上定义的方法
    // 下面要在非结构体上定义方法
    aaa := MyInt(5)
    bbb := MyInt(10)
    fmt.Println(aaa.MyMethod(bbb))
    
}