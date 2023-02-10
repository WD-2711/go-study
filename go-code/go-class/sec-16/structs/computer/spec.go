// spec.go
// 创建导出结构体Spec，首字母要大写哦，你知道为什么

package computer

type Spec struct {
    Maker string // 可以导出
    model string // 不可以导出
    Price int   // 可以导出
}

