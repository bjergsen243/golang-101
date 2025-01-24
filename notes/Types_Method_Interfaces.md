# Types, Methods, and Interfaces

## Types in Go

- Abstract type là chỉ định type đó nên làm gì (what), chứ không phải cách nó làm (how)
- Concrete type chỉ định cả what và how

## Methods

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

## Pointer Receivers and Value Receivers

- Khi 1 type có bất kỳ pointer receiver method, cần phải sử dụng pointer receiver cho tất cả các method, kể cả cái mà không làm thay đổi receiver
- Nếu method có 1 pointer receiver và sẽ không hoạt động cho 1 nil receiver, kiểm tra nil và trả về lỗi

## Methods Are Functions Too

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

## Functions Versus Methods
