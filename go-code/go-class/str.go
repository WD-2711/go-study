package main

import (
    "fmt"
    "unicode/utf8"
)


func main(){
    str := "hello"
    for _,ss := range str{
        fmt.Printf("%c ", ss)
    }
    fmt.Printf("\n")
    
    // 比较特殊的字符串，此时输出会有问题，原因就是ñ在utf-8编码中占用了c3与b1两个字节，所以这样输出就会出错
    str = "Señor"
    for i := 0; i < len(str); i++{
        fmt.Printf("%x %c\n", str[i], str[i])
    }
    fmt.Printf("\n")
    
    //用rune可以解决上述问题，rune是int32的别称，可以用rune来打印字符
    runes := []rune(str)
    for i := 0; i < len(runes); i++{
        fmt.Printf("%c %x\n", runes[i], runes[i])
    }
    
    // 直接使用for 循环更简单
    for i, v := range str{
        fmt.Printf("%d %c\n",i, v)
    }
    
    // 直接用字节切片构造字符串
    byteSlice := []byte{0x43, 0x61, 0x66, 0xc3, 0xa9}
    str = string(byteSlice)
    fmt.Println(str)
    
    // 计算字符串的长度
    word1 := "Señor"
    word2 := "xcl11"
    fmt.Printf("%d %d\n", utf8.RuneCountInString(word1), utf8.RuneCountInString(word2))
    
    // 字符串是不可变的，如果想修改，就先把字符串转化为rune切片，对切片进行修改后再转化为字符串
    str2 := "hello"
    // str2[0] = "x"
    fmt.Println(str2)
    nstr2 := []rune(str2) // 将字符串转化为rune切片
    fmt.Println(nstr2)
    nstr2[0] = 'x' // 不能用双引号，要用单引号
    nstr3 := string(nstr2) // 将rune切片转化为字符串
    fmt.Println(nstr3)
}