package main

import "fmt"

/* 简单的示例 */
type Income interface {
    calculate() int
    source() string
}
// 收益1
type FixedBilling struct {
    projectName string
    biddenAmount int
}
// 收益2
type TimeAndMeterial struct {
    projectName string
    noOfHours int
    hourlyRate int
}

func (fb FixedBilling) calculate() int {
    return fb.biddenAmount
}
func (fb FixedBilling) source() string {
    return fb.projectName
}

func (tm TimeAndMeterial) calculate() int {
    return tm.noOfHours * tm.hourlyRate
}
func (tm TimeAndMeterial) source() string {
    return tm.projectName
}

func calculateNetIncome(ic []Income) {
    var netincome int = 0
    for _, income := range ic {
        fmt.Printf("Income from %s = %d\n", income.source(), income.calculate())
        netincome += income.calculate()
    }
    fmt.Printf("Net income of organisation = %d\n", netincome)
}

/* 新增收益流 */

func main() {
    /* 简单的多态示例，感觉就是多个类型都实现了这个接口 */
    p1 := FixedBilling{projectName: "p1", biddenAmount: 5000}
    p2 := FixedBilling{projectName: "p2", biddenAmount: 1000}
    p3 := TimeAndMeterial{projectName: "p3", noOfHours: 160, hourlyRate: 25}
    incomes := []Income{p1, p2, p3}
    calculateNetIncome(incomes)
    
    /* 新增收益流，就是新加一个赚钱的方式，省略了，没意思 */
}