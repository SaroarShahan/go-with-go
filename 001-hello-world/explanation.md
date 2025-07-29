## 🧑🏼‍💻 `package main`

- Think of `package` as a way to organize our code.
- Every Go file **must belong to a package**.

- When we write a Go program that we want to **run**, we use `package main`.

- It tells Go: 👉 “This is a standalone program, and it should start from the `main()` function.”

## 🧑🏼‍💻 `import "fmt"`

- `import` means we're **bringing in a built-in library** that gives us extra powers.
- `"fmt"` is a package in Go that stands for **"format"**.
- It helps us **print messages** to the screen, like:

```go
fmt.Println("Hello, World!")
```
- So, when we write:
```go
import "fmt"
```
- We're saying: 👉 “We want to use the `fmt` package to print stuff!”