package main

import "oop/employee"

func main(){
//     e := employee.Employee {
//         FirstName: "clong",
//         LastName: "xie",
//         TotalLeaves: 30, 
//         LeavesTaken: 20,
//     }
    
    /* 
    以下情况会直接输出 has 0 leaves remaining
    这是不太优雅的，因为0值没有什么卵用，如果某类型的0值不可用，需要我们来隐藏这个类型，以避免从其他包中直接访问。
    java是用构造器来解决上述问题的。
    那Go呢，它就甘心被java比下去了吗？？
    */
//     var e employee.Employee
//     e.LeavesRemaining()
    
    
    /* 
    当然不！！Go使用newT的函数来解决上述问题，当然newT函数是定义在oop/employee/employee.go里面咯 
    */
    e := employee.New("clong", "xie", 30, 20)
    e.LeavesRemaining()
                      
}