# Learn golang with love
# 1. Variables / Constant / Declaration

> *Một biến khi được khai báo thì luôn luôn phải sử dụng ít nhất một lần.*
>
- Khai báo biến trong Golang

```go
package main

import "fmt"

func main() {
	var a = "string" // Using var keyword
	var b, c int = 2, 34 // 2 var with var keyword
	var d int // no init value
	e := true // Golang auto inffer this value to data type 

	fmt.Println(a, b, c, d, e)
}
```

- Hàng trong golang: Sử dụng từ khóa const

```go
const a = 10

func main() {
	fmt.Println(a) // co the khong can dung toi hang so, van khong bao loi
}
```

- Global scope: Chúng ta có thể tạo ra một nhóm những biến có scope là global với từ khóa var nằm bên ngoài hàm main

```go
var (
	a = 10
	b = "string"
)

func main() {
	fmt.Print(a, b)
}
```

- Ép kiểu

```go
func main() {
	var i int = 1
	var b float32 = float32(i)
	var s string = strconv.Itoa(i)
	fmt.Println(i, b, s)
}
```