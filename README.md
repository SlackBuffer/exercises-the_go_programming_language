- *Parameters* are local variables whose value or *arguments* are supplied by the caller
- An *object* is simply a value or variable that has methods
- 调用方法时，编译器会自动为 receiver 做恰当的值-指针类型的隐式转换，所以可以怎么写简洁怎么来（receiver argument 和 receiver parameter 不一致，分别为 `*T` 和 `T` 中的任意一个）
- Like maps (`make`, literal) and slices(`make`, literal), channels (`make`) must be created before use
    - [x] pointers (`new(T)` returns a `*T`; using `&`)
    - Functions are also reference types
- printf
	
    ```go
    package fmt
    func formatOneValue(x interface{}) string {
        if err, ok := x.(error); ok {
            return err.Error()
        }
        if str, ok := x.(Stringer); ok {
            return str.String()
        }
        // ...all other types...
    }
    ```

- Dereference: loads the value
---
- [ ] https://golang.org/doc/effective_go.html
- [ ] https://golang.org/doc/code.html
- Review (32)
    - [x] basic-structure
    - [x] assignments
    - [x] declarations-initializations 
    - [x] pointers
    - [x] control-flow
    - [x] visibility-scope-lifetime
    - [x] strings_Unicode_UTF-8_byte-slice
    - [x] packages
    - [x] integers
    - [x] floating-point
    - [x] booleans
    - [x] complex-numbers
    - [x] operators
    - [x] constants
    - [x] web
    - [x] verb-adv
    - [x] io
    - [x] cmd
    - [x] templates
    - [x] idiomatic
    - [x] go-toolchain
    - [x] array
    - [x] slice
    - [x] maps
    - [x] structs
    - [x] JSON
    - [x] functions
    - [x] methods
    - [x] interfaces
    - [x] goroutine-channel
    - [x] concurrency-shared_variables
    - [ ] testing


Continues at p348