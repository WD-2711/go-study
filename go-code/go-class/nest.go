package main

import "fmt"

/* 结构体1及其方法1 */ 
type author struct {
    firstName string
    lastName string
    bio string
}

func (a author) fullName() string {
    // fmt.Sprintf会将格式化后的字符串保存到一个变量中，而fmt.Printf只是简单的输出到屏幕中
    return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

/* 结构体2及其方法2 （可嵌套）*/
type post struct {
    title string
    content string
    author
}
func (p post) details() {
    fmt.Println("Title: ", p.title)
    fmt.Println("Content: ", p.content)
    // fmt.Println("Author: ", p.author.fullName())
    // 可以写为：
    fmt.Println("Author: ", p.fullName())
    fmt.Println("Bio: ", p.author.bio)
}

/* 结构体切片的一个小示例 */
type website struct {
    posts []post
}
func (w website) contents(){
    fmt.Println("Content of Website\n")
    for _, v := range w.posts {
        v.details()
        fmt.Println()
    }
}

func main() {
    /* 结构体嵌套的一个小示例 */
    a1 := author {
        "clong",
        "xie",
        "genius",
    }
    p1 := post {
        "xie's life",
        "content",
        a1,
    }
    p1.details()
    
    /* 结构体切片的嵌套 */
    a2 := author {
        "wjh",
        "jinhe",
        "good",
    }
    p2 := post {
        "jinhe's life",
        "content",
        a2,
    }
    a3 := author {
        "zhamin",
        "qin",
        "good",
    }
    p3 := post {
        "zhamin's life",
        "content",
        a3,
    }
    w := website {
        posts: []post{p1, p2, p3},
    }
    w.contents()
}