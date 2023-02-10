package main

import "fmt"


func printarr(a [3][2]string){
    for _, v1 := range a{
        for _, v2 := range v1{
            fmt.Printf("%s ", v2)
        }
        fmt.Printf("\n")
    }
}


func slice_sub2(numbers []int){
    for i, _ := range numbers {
        numbers[i] -= 2
    }
}

func arr_sub2(numbers [3]int){
    for i, _ := range numbers {
        numbers[i] -= 2
    }
}


func main() {
    // 长度为3的int数组
    var a [3]int
    a[0] = 13
    a[1] = 14
    a[2] = 15
    fmt.Println(a)
    
    // 另一种写法
    b := [3]int{1, 2, 3}
    fmt.Println(b)
    
    // 只是对第一个元素赋值为12
    c := [3]int{12}
    fmt.Println(c)
    
    // 忽略数组长度，len()数组长度
    d := [...]int{9, 9, 9, 9, 9}
    fmt.Println(d)
    fmt.Println(len(d))
    
    // 遍历数组
    aa := [...]float64{1.1, 2.2, 3.3}
    sum := float64(0) // 强制类型转换
    // i 代表索引，v代表具体值
    for i, v := range aa {
        fmt.Printf("%d the element of aa is %.2f\n", i, v)
        sum += v
    }
    fmt.Println("\n sum is ", sum)
    
    
    // 多维数组
    m := [3][2]string{
        {"lion", "bbb"},
        {"ccc", "ddd"},
        {"eee", "fff"},//这个逗号是必须的哈，要不就给你加；了
    }
    printarr(m)
    
    
    // 创建切片，感觉和python差不多
    // 但是对切片的修改会反映到底层数组中
    // 切片的长度指的是切片中的元素数，切片的容量指的是以切片的开头为开头，到底层数组末尾中的元素数
    // 切片也可以重置长度与容量
    s := [5]int{3, 4, 5, 6, 7}
    var bb []int = s[1:4]
    fmt.Println(bb, s)
    for i, _ := range bb{
        bb[i] = 0
    }
    fmt.Println(bb, s)
    fmt.Printf("len %d, cap %d\n", len(bb), cap(bb))
    bb = s[:1]
    fmt.Printf("len %d, cap %d\n", len(bb), cap(bb))
    

    // make 创建切片 make([]T, len, cap)
    ii := make([]int, 5, 5)
    fmt.Println(ii)
    
    // 给切片append看长度与容量
    cars := []string{"11", "22", "33"}
    // 注意cars := [...]string{"11", "22", "33"}的话就是数组
    fmt.Printf("len %d, cap %d\n", len(cars), cap(cars))
    cars = append(cars, "44")
    fmt.Printf("len %d, cap %d\n", len(cars), cap(cars))
    // 可以看到，添加了一个元素，长度加1，容量*2
    
    // 切片类型的零值（啥也没有）是nil，nil切片的长度和容量都是0
    var names []string
    if names == nil {
        names = append(names, "bob", "alice")
        fmt.Println(names)
        fmt.Printf("len %d, cap %d\n", len(names), cap(names))
    }
    
    // 可以使用...运算符来连接两个切片
    slice1 := []string{"11", "22", "3"}
    slice2 := []string{"5", "6"}
    slice3 := append(slice1, slice2...)
    fmt.Println(slice3)
    
    // 数组的传递，如果直接传递给函数的话，在函数内对数组的修改在函数外面是不起作用的。
    // 切片的传递与数组不同，在函数内的修改在函数外面是可以体现的。
    nos := []int{6, 7, 8}
    slice_sub2(nos)
    // nos := [...]int{6, 7, 8}
    // arr_sub2(nos)
    fmt.Println(nos)
    
    
    // 多维切片
    pls := [][]string {
        {"C", "C++"},
        {"JAVA"},
        {"PYTHON", "PHP"},
    }
    for _, v1 := range pls {
        for _, v2 := range v1{
            fmt.Printf("%s ", v2)
        
        }
        fmt.Printf("\n")
    }
    
    // 由于切片是对数组的引用，因此只要切片在内存中，那么数组就不能被回收，数组就会一直在内存中占用内存空间
    // 可以使用copy(dst, src []T)来对切片进行复制，使用新的切片，这样的话旧的切片引用的数组就会被回收
    bigArr := [...]string{"usa", "china", "russia", "korea", "japan"}
    slice11 := bigArr[3:]
    newslice := make([]string, len(slice11))
    copy(newslice, slice11)
    fmt.Println(newslice)
    
    
    
    
}