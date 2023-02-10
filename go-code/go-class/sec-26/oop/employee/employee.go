package employee

import "fmt"

/* 不太优雅的写法 */
// 一个结构体，绑定了一个方法，就和类很相似有木有?
// type Employee struct {
//     FirstName string // 首字母必须大写才能访问到
//     LastName string
//     TotalLeaves int
//     LeavesTaken int
// }

// func (e Employee) LeavesRemaining() {
//     fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
// }


/* 更nb的写法，使用了Go的newT函数 */
// 当然，首字母全变小写，让别人访问不到嘿嘿嘿
type employee struct {
    firstName string
    lastName string
    totalLeaves int
    leavesTaken int
}

// 只有New函数可以被外面访问到哦，因为首字母大写
func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {
    e := employee {firstName, lastName, totalLeave, leavesTaken}
    return e
}

func (e employee) LeavesRemaining(){
    fmt.Printf("%s %s has %d leaves remaining\n", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}