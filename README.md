# Learn golang with love

# 2. ****Data Types****

### Boolean

- true
- false

### Numerics

- Integer

    ```go
    func main() {
    	var i int = -1
    	var u uint8 = 12 // u must be > 0
    	fmt.Println(i, u)
    }
    ```

    - signed integer: int8, int16, int32, int 64
    - unsigned integer: uint8, uint16, unit32, uint64
- Float:

    ```go
    func main() {
    	var i float32 = 13.2
    	var e float32 = 2.1e2
    	fmt.Println(i, e)
    }
    ```

- Complex: số phức

    ```go
    func main() {
    	var i complex64 = 13.2 + 32i
    	fmt.Println(i)
    	fmt.Println(real(i)) // Phần thực
    	fmt.Println(imag(i)) // Phần ảo
    }
    ```


### Text

- String
    - Parse String to another data type

    ```go
    func main() {
    	b, _ := strconv.ParseBool("true")
    	f, _ := strconv.ParseFloat("3.1415", 64)
    	i, _ := strconv.ParseInt("-42", 10, 64)
    	u, _ := strconv.ParseUint("42", 10, 64)
    
    	fmt.Println(b, f, i, u)
    }
    ```

    - Substring / split / …. using strings package

    ```go
    func main() {
    	var s string = "1234"
    	var sub string = s[0:3] // sub string from index 0 and get 3 items
    	var isContain = strings.Contains(s, "1")
    	spiltString := strings.Split(s, "") // [1 2 3 4]
    }
    ```

- Byte

    ```go
    func main() {
    	var s string = "1234"
    	var b []byte = []byte(s) // [49 50 51 52]
    }
    ```

- Rune