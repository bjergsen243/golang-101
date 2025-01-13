# Golang 101

Go là ngôn ngữ static type. Nghĩa là khi biên dịch thì Go cần biết kiểu dữ liệu của các biến và tham số

-> Nếu dùng generics sẽ phá vỡ nguyên tắc này. Nghĩa là khi tới runtime thì mới check được

## Tài liệu tham khảo

- **Sách**: _Learning Go_
- [Go By Example](https://gobyexample.com/)
- [200Lab Blog: Golang](https://200lab.io/blog/tag/golang/)

---

## Table of contents

| Title | Description |
|-------|-------------|
| [Golang 101](#golang-101) | Introduction to Golang and its static type nature. |
| [Tài liệu tham khảo](#tài-liệu-tham-khảo) | References for learning Golang. |
| [Table of contents](#table-of-contents) | Overview of the document structure. |
| [Kiểu dữ liệu và Khai báo trong Golang](#kiểu-dữ-liệu-và-khai-báo-trong-golang) | Data types and variable declarations in Golang. |
| [Mảng và Slice trong Go](#mảng-và-slice-trong-go) | Arrays and slices in Golang. |
| [Strings và Runes và Bytes](#strings-và-runes-và-bytes) | Working with strings, runes, and bytes in Golang. |
| [Maps](#maps) | Using maps in Golang. |
| [Struct](#struct) | Structs in Golang and their usage. |
| [for, 4 cách](#for-4-cách) | Different ways to use for loops in Golang. |
| [Functions](#functions) | Functions in Golang, including variadic parameters and anonymous functions. |
| [Pointers - Con trỏ](#pointers---con-trỏ) | Understanding pointers in Golang. |
| [Types, Methods, and Interfaces](#types-methods-and-interfaces) | Types, methods, and interfaces in Golang. |

---

## Types, Methods, and Interfaces

-

### Types in Go

- Abstract type là chỉ định type đó nên làm gì (what), chứ không phải cách nó làm (how)
- Concrete type chỉ định cả what và howÏ

### Methods

- Khai báo method giống khai báo function, thêm 1 điều kiện: receiver.
- Receiver ở giữa keywork `func` và tên của method. Tên của receiver ở trước type. Thường là viết tắt của type's name, thường là chữ cái đầu
- Ví dụ

  ```go
  type Person struct {
    FirstName string
    LastName string
    Age int
  }
  func (p Person) String() string {
    return “fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)”
  }
  ```

### Pointer Receivers and Value Receivers

- Khi 1 type có bất kỳ pointer receiver method, cần phải sử dụng pointer receiver cho tất cả các method, kể cả cái mà không làm thay đổi receiver
- Nếu method có 1 pointer receiver và sẽ không hoạt động cho 1 nil receiver, kiểm tra nil và trả về lỗi

### Methods Are Functions Too

- Có thể sử dụng method thay function
- Ví dụ:

  ```go
  type Adder struct {
    start int
  }

  func (a Adder) AddTo(val int) int {
    return a.start + val
  }
  ```

- Từ đây có các cách sử dụng sau:

  - Thông thường:

    ```go
    myAdder := Adder{start: 10}
    fmt.Println(myAdder.AddTo(5)) // prints 15
    ```

  - Method value:

    ```go
    f1 := myAdder.AddTo
    fmt.Println(f1(10)) // prints 20
    ```

  - Method expression:

    ```go
    f1 := Adder.AddTo
    fmt.Println(f2(myAdder, 15)) // prints 25
    ```

### Functions Versus Methods
