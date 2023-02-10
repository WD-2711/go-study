package main

import (
    "fmt"
)

// 一个接口的定义
type VowelsFinder interface {
    // 定义了 FindVowels 方法，方法的返回值为rune的切片
    FindVowels() []rune
}

type MyString string

// MyString类型实现了FindVowels方法
// 此方法的作用就是找ms内所有的元音
func (ms MyString) FindVowels() []rune {
    var vowels []rune
    // 遍历 ms 字符串
    for _, rune := range ms {
        if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u'{
            vowels = append(vowels, rune)
        }
    }
    return vowels
} 


/*接口的实际应用场景*/
type SalaryCalculator interface {
    CalculateSalary() int
}
// 终身雇佣制的员工
type Permanent struct {
    empId int
    basicpay int
    pf int // 感觉可能是福利
}
// 签合同的员工
type Contract struct {
    empId int
    basicpay int
}
func (p Permanent) CalculateSalary() int {
    return p.basicpay + p.pf
}
func (c Contract) CalculateSalary() int {
    return c.basicpay
}
func totalExpense(s []SalaryCalculator) {
    expense := 0
    for _, v := range s{
        expense += v.CalculateSalary()
    }
    fmt.Printf("%d\n", expense)
}


/*接口的内部表示*/
type Test interface {
    Tester()
}
type MyFloat float64
func (m MyFloat) Tester(){
    fmt.Println(m)
}
func describe(t Test){
    fmt.Printf("interface type %T value %v\n", t, t)
}

/*空接口*/
func describe11(i interface{}){
    fmt.Printf("type %T, value %v\n", i, i)
}

/*类型断言*/
func assert(i interface{}){
    s, ok := i.(int) // 获得接口的底层int值，如果不加ok可能会报错
    fmt.Println(s, ok) 
}

/*类型选择*/
func findType(i interface{}) {
    switch i.(type) {
        case string:
            fmt.Printf("I am a string and my value is %s\n", i.(string))
        case int:
            fmt.Printf("I am a int and my value is %d\n", i.(int))
        default:
            fmt.Printf("Unknown type\n")
    }
}


/*类型与接口相比较*/
type Describer22 interface{
    Describe22()
}
type Person22 struct {
    name string
    age int
}
func (p Person22) Describe22() {
    fmt.Printf("%s %d\n", p.name, p.age)
}
func findType22(i interface{}){
    switch v := i.(type) {
        case Describer22: // 实现了Describer22方法的类型Person22
            v.Describe22()
        default:
            fmt.Printf("unknown type\n")
    }
}

func main(){
    name := MyString("xcllove someone")
    
    // 定义了一个接口
    var v VowelsFinder
    
    // 由于name是MyString类型，MyString类型又实现了FindVowels方法，因此可以直接将name赋给接口v
    v = name
    
    // 这里神奇的是%c直接能输出切片的所有元素
    fmt.Printf("%c",v.FindVowels())
    fmt.Printf("\n")
    
    // 接口应该怎么用
    // totalExpense可以扩展新的员工类型，而不需要修改任何代码，有了新的员工类型之后不用修改totalExpense函数，只需要实现SalaryCalculator接口即可
    pemp1 := Permanent{1, 5000, 20}
    pemp2 := Permanent{2, 6000, 30}
    cemp1 := Contract{3, 3000}
    employees := []SalaryCalculator{pemp1, pemp2, cemp1}
    totalExpense(employees)
    
    // 接口的内部表示
    // 数据f赋给接口t，这样接口t的类型为main.MyFloat，接口的值为89.7
    var t Test
    f := MyFloat(89.7)
    t = f
    describe(t)
    t.Tester()
    
    // 空接口
    // 没有包含方法的接口为空接口，默认所有类型都实现了空接口
    s := "helloworld"
    describe11(s)
    i := 55
    describe11(i)
    strt := struct {
        name string
    }{
        name: "Naveen R",
    }
    describe11(strt)
    var sss interface{} = 56
    describe11(sss)
    
    // 类型断言， 用于提取接口的底层值
    var s1 interface{} = "hello"
    assert(s1)
    
    // 类型选择 type switch
    // 与传统的switch类似，只不过传统的指定的是值，type switch指定的是类型
    findType("Naveen")
    findType(77)
    findType(9.1)
    
    
    // 一个类型与接口相比较。
    // 如果一个类型实现了接口，那么该类型与其实现的接口就可以互相比较
    fmt.Println("--------------")
    findType22("Naveen")
    p := Person22{
        name: "naveen R",
        age: 25,
    }
    findType22(p)
    
    
}

