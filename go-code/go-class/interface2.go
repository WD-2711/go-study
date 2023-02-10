package main

import "fmt"

/* 指针接收器与值接收器进一步详解 */
type Describer interface {
    Describe()
}
type Address struct {
    state string
    country string
}
func (a *Address) Describe() {
    fmt.Printf("state %s country %s\n", a.state, a.country)
}
type Person struct {
    name string
    age int
}
func (p Person) Describe() {
    fmt.Printf("name %s age %d\n", p.name, p.age)
}

/* 实现多个接口 */
type SalaryCalculator interface {
    DisplaySalary()
}
type LeaveCalculator interface {
    CalculateLeavesLeft() int
}
type Employee struct {
    firstName string
    lastName string
    basicPay int
    pf int
    totalLeaves int
    leavesTaken int
}
func (e Employee) DisplaySalary() {
    fmt.Printf("%s %s has salary %d\n", e.firstName, e.lastName, (e.basicPay+e.pf))
}
func (e Employee) CalculateLeavesLeft() int {
    return (e.totalLeaves - e.leavesTaken)
}

/* 接口的嵌套 */
// 加上个 EmployeeOperation 接口，直接嵌套前面的 SalaryCalculator 与 LeaveCalculator 接口
type EmployeeOperation interface {
    SalaryCalculator
    LeaveCalculator
}

/* 接口的零值 */
type nilIf interface {
    DisplaySalary()
}

func main() {
    var d1 Describer
    // 用值作接收器既可以用值来调用，也可以用指针来调用
    p1 := Person{"sam", 25}
    d1 = p1
    d1.Describe()
    
    p2 := Person{"James", 22}
    d1 = &p2
    d1.Describe()
    
    var d2 Describer
    a := Address{"shandong", "china"}
    // 会出错，因为这里只接收指针类型，不能接受值类型
    // d2 = a
    d2 = &a
    d2.Describe()
    
    
    /* 实现多个接口 */
    e := Employee {
        firstName: "clong",
        lastName: "xie",
        basicPay: 666666,
        pf: 200,
        totalLeaves: 30,
        leavesTaken: 5,
    }
    var s SalaryCalculator
    s = e
    s.DisplaySalary()
    var l LeaveCalculator
    l = e
    fmt.Printf("\nleavesLeft %d\n", l.CalculateLeavesLeft())
    
    /* 接口的嵌套 */
    var newIf EmployeeOperation = e
    newIf.DisplaySalary()
    fmt.Println(newIf.CalculateLeavesLeft())
    
    /* 接口的零值 */
    var if1 nilIf
    if if1 == nil {
        fmt.Println("Interface is nil")
    }
}
