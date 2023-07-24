# Learn golang with love

# 6. ****Function****

### **Variadic Functions**

```jsx
func add(args ...int) int {
  total := 0
  for _, v := range args {
    total += v
  }
  return total
}
func main() {
  fmt.Println(add(1,2,3))
}
```

### **Closure**

```jsx
func main() {
  add := func(x, y int) int {
    return x + y
  }
  fmt.Println(add(1,1))
}
```

### **Returning Multiple Values**

```jsx
func f() (int, int) {
  return 5, 6
}

func main() {
  x, y := f()
}
```
