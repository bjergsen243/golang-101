# Golang 101

Go là ngôn ngữ static type. Nghĩa là khi biên dịch thì Go cần biết kiểu dữ liệu của các biến và tham số

-> Nếu dùng generics sẽ phá vỡ nguyên tắc này. Nghĩa là khi tới runtime thì mới check được

## Tài liệu tham khảo

- **Sách**: _Learning Go_
- [Go By Example](https://gobyexample.com/)
- [200Lab Blog: Golang](https://200lab.io/blog/tag/golang/)

---

## Công cụ hữu ích

### Lint

- Dùng để kiểm tra chất lượng mã nguồn, phát hiện các lỗi code style hoặc tiềm ẩn.

- Cài đặt

  ```bash
  go install golang.org/x/lint/golint@latest
  ```

- Sử dụng

  ```bash
  golint ./...
  ```

---

### Kiểm tra lỗi

- Kiểm tra lỗi cơ bản:

  - Sử dụng lệnh go vet để phát hiện các vấn đề tiềm ẩn trong mã nguồn:

  ```bash
  go vet ./...
  ```

- Sử dụng golangci-lint:
  - Công cụ này kết hợp nhiều linter, bao gồm cả golint và go vet.
  - Cấu hình và hướng dẫn sử dụng: [golangci-lint](https://golangci-lint.run/usage/configuration)

👉 Khuyến nghị:

- Sử dụng golint và go vet cho các kiểm tra thông thường.
- Chỉ dùng golangci-lint nếu cả team đồng ý để đảm bảo tính thống nhất.

---

### Makefile

#### Thành phần

1. **Target:**

   - Mỗi lệnh trong Makefile được gọi là một **target**.
   - Target nằm ở trước dấu hai chấm `:`.
   - Các thành phần sau dấu `:` là các target phụ (dependencies) cần được thực thi trước target chính.

   Ví dụ:

   ```make
   build: clean
       go build .
   ```

   - Trong ví dụ trên, target `build` phụ thuộc vào target `clean`. Trước khi `build` chạy, target `clean` sẽ được thực thi.

2. **DEFAULT_GOAL:**

   - `DEFAULT_GOAL` là target mặc định sẽ chạy nếu bạn không chỉ định target cụ thể.
   - Ví dụ:

     ```make
     .DEFAULT_GOAL := build
     ```

3. **`.PHONY`:**

   - Dùng để khai báo các target không tương ứng với tệp tin hoặc thư mục thực tế.
   - Điều này giúp Makefile tránh nhầm lẫn nếu trong dự án có thư mục hoặc tệp có tên trùng với target.

   Ví dụ:

   ```make
   .PHONY: clean
   clean:
       rm -rf build/
   ```

#### Lợi ích của Makefile

- Tự động hóa các lệnh phức tạp.
- Đảm bảo tính nhất quán khi chạy lệnh trên các máy khác nhau.
- Giúp quản lý quy trình build và kiểm tra mã nguồn dễ dàng hơn.

---

## Kiểu dữ liệu và Khai báo trong Golang

### Kiểu dữ liệu Built-in

Go cung cấp một số kiểu dữ liệu cơ bản (built-in types) mà bạn có thể sử dụng trong mọi chương trình. Các kiểu dữ liệu này bao gồm:

- **Boolean**: Kiểu dữ liệu này chỉ có hai giá trị là `true` hoặc `false`.
- **Integer**: Dùng để biểu diễn các số nguyên. Ví dụ: `int`, `int32`, `int64`.
- **Float**: Dùng để biểu diễn các số thập phân. Ví dụ: `float32`, `float64`.
- **String**: Dùng để biểu diễn chuỗi văn bản.

#### The Zero Value (Giá trị mặc định)

Khi bạn khai báo một biến mà không gán giá trị, Go sẽ tự động gán cho biến đó một giá trị mặc định (zero value). Ví dụ:

- Các kiểu số (integer, float) sẽ được gán giá trị `0`.
- Các kiểu Boolean sẽ được gán giá trị `false`.
- Các kiểu chuỗi (`string`) sẽ được gán giá trị chuỗi rỗng `""`.

#### Literals (Chữ số nguyên thủy)

`Literal` là các giá trị được viết trực tiếp trong mã nguồn. Có thể là số, ký tự hoặc chuỗi.

- **Số nguyên**:

  - Bạn có thể khai báo như sau:

    ```go
    1, 1_234, 0b1010, 0o12, 0xF
    ```

  - `0b` biểu thị số nhị phân, `0o` biểu thị số hệ bát phân (octal), `0x` biểu thị số hệ thập lục phân (hexadecimal). Tuy nhiên, hãy tránh sử dụng chúng vì chúng có thể gây nhầm lẫn.

- **Số thập phân (Float)**:

  - Dùng dấu chấm (.) để biểu thị giá trị số thập phân.

    ```go
    6.03e23  // 6.023^23
    ```

  - Go có hai kiểu `float32` và `float64` cho phép bạn chọn độ chính xác.

- **Rune literal**:

  - Rune là một kiểu dữ liệu đặc biệt để biểu thị một ký tự. Rune được bao quanh bởi dấu nháy đơn (`'`), ví dụ:

    ```go
    'a', '1', '', ' '
    ```

**Tóm lại**:

- Sử dụng hệ cơ số 10 để biểu diễn số để tránh nhầm lẫn, đặc biệt là khi sử dụng số thập lục phân (hexadecimal) trong biểu thức ký tự (rune literal).
- Sử dụng dấu ngoặc kép (`""`) để tạo chuỗi (string).
- Nếu bạn cần sử dụng dấu backslash (`\`), dấu nháy kép (`"`) hoặc xuống dòng trong chuỗi, hãy sử dụng raw string literal (bao quanh bởi dấu backtick ``).

#### Các kiểu số học (Numeric Types)

- **Byte**: Tương đương với `uint8`. Dùng để biểu thị một giá trị số không âm trong phạm vi 0-255.
- **Rune**: Tương đương với `int32`. Dùng để biểu thị một ký tự trong bảng mã Unicode.
- **uintptr**: Dùng để lưu trữ địa chỉ bộ nhớ.

### Khai báo biến trong Golang

Khi khai báo biến, bạn cần phải lưu ý những điểm sau:

- **Tránh sử dụng `:=` khi khai báo biến với giá trị mặc định (zero value)**: Khi khai báo biến với giá trị mặc định, hãy sử dụng từ khóa `var` thay vì `:=`.

  ```go
  var x int  // x sẽ có giá trị mặc định là 0
  ```

- **Khi gán giá trị cho một hằng số (constant) hoặc một literal với kiểu dữ liệu chưa xác định**, bạn nên sử dụng cú pháp khai báo rõ ràng với từ khóa `var`. Ví dụ:

  ```go
  var x byte = 20
  ```

- **Tránh khai báo biến trong phạm vi toàn cục (global scope)** vì nó có thể làm cho chương trình phức tạp và khó kiểm soát. Chỉ nên khai báo các biến trong phạm vi local (trong các function).

- **Các biến local cần được sử dụng (read)**, tránh khai báo các biến không cần thiết, bởi chúng có thể gây rối cho việc phân tích mã sau này.

- **Hằng số (`const`) có thể khai báo mà không cần phải đọc chúng**, bởi vì giá trị của chúng đã được tính toán trong quá trình biên dịch (compile time). Nếu hằng số không được sử dụng trong chương trình, nó sẽ không ảnh hưởng đến quá trình biên dịch và không được đưa vào file thực thi cuối cùng.

**Tóm lại**:

- Hãy khai báo biến với `var` khi bạn muốn có kiểu dữ liệu rõ ràng, đặc biệt khi bạn làm việc với giá trị mặc định.
- Tránh sử dụng `:=` nếu bạn muốn khai báo một hằng số hoặc một giá trị với kiểu dữ liệu không xác định.
- Hãy đảm bảo rằng mọi biến được khai báo đều được sử dụng trong mã của bạn để tránh việc khai báo thừa.

---

## Mảng và Slice trong Go

Mảng là kiểu dữ liệu khá đặc biệt trong Go, và ít khi được sử dụng trực tiếp vì những lý do sau:

- **Kích thước mảng là một phần của kiểu dữ liệu**: Trong Go, mảng không chỉ được xác định bởi kiểu dữ liệu mà còn bởi kích thước (size) của mảng. Ví dụ, một mảng có kích thước 3 (`[3]int`) là khác so với một mảng có kích thước 4 (`[4]int`). Điều này có thể gây khó khăn khi làm việc với các mảng có kích thước không cố định.
- **Không thể sử dụng biến làm kích thước mảng**: Go yêu cầu kích thước mảng phải được xác định tại thời điểm biên dịch (compile time), chứ không phải runtime. Do đó, bạn không thể sử dụng biến làm kích thước cho mảng trong Go.

- **Không thể ép kiểu mảng**: Một mảng trong Go không thể được ép kiểu thành một mảng có kiểu khác, và bạn không thể viết một hàm xử lý mảng có kích thước bất kỳ. Bạn cũng không thể gán một mảng có kích thước khác cho mảng khác.

- **Chỉ hỗ trợ mảng 1 chiều**: Go chỉ hỗ trợ mảng một chiều, nhưng bạn có thể tạo mảng nhiều chiều. Ví dụ: `var x [2][3]int`.

> **Tóm lại về mảng**:

- Mảng thích hợp khi bạn biết trước kích thước (length) của mảng và kích thước này không thay đổi.
- Mảng tồn tại chủ yếu để hỗ trợ cho kiểu dữ liệu `slice`.

### Slices trong Go

Slices là một trong những kiểu dữ liệu mạnh mẽ và linh hoạt nhất trong Go. So với mảng, slices có nhiều ưu điểm, và là lựa chọn phổ biến trong Go khi làm việc với dữ liệu tuần tự.

#### Khai báo Slice

- **Slice literal**: Bạn có thể khai báo slice bằng cách sử dụng cú pháp slice literal. Ví dụ: `var x = []int{10, 20, 30}`.
- **Nil zero value**: Giá trị mặc định (zero value) của slice là `nil`. Bạn có thể khai báo một slice rỗng như sau: `var x []int`.
- **`make` function**: Hàm `make` được sử dụng để tạo một slice với độ dài và capacity được chỉ định. Ví dụ: `x := make([]int, 5, 10)` tạo ra một slice có độ dài 5 và capacity 10.

#### Đặc điểm của Slice

- **Không có kích thước cố định**: Slice trong Go không cần chỉ định kích thước (length) khi khai báo. Điều này làm cho slice trở thành lựa chọn tốt hơn khi bạn không biết trước số lượng phần tử.
- **Dễ dàng thay đổi kích thước**: Một điểm mạnh của slice là khi bạn append thêm phần tử vào slice, Go sẽ tự động điều chỉnh kích thước của slice nếu cần. Tuy nhiên, điều này cũng có thể dẫn đến panic nếu bạn cố gắng thêm phần tử vào một slice có capacity nhỏ hơn length.

- **Không thể so sánh slices**: Bạn không thể so sánh hai slice trực tiếp trong Go (trừ khi so sánh với `nil`). Tuy nhiên, bạn có thể sử dụng hàm `DeepEqual` từ package `reflect` để so sánh nội dung của hai slice.

- **Chia sẻ vùng nhớ**: Nếu bạn tạo một slice mới từ một slice khác (subslice), hai slice này sẽ chia sẻ cùng một vùng nhớ. Điều này có nghĩa là nếu bạn thay đổi giá trị trong một slice, các slice khác có thể bị ảnh hưởng.

#### Capacity của Slice

- **Capacity**: Mỗi slice có một capacity, tức là dung lượng bộ nhớ tối đa mà slice có thể chứa. Capacity có thể lớn hơn length. Khi bạn thêm phần tử vào slice và length vượt quá capacity, Go sẽ tự động cấp phát một slice mới với capacity lớn hơn và sao chép dữ liệu từ slice cũ vào slice mới.

- **Tăng capacity tự động**: Từ Go 1.14, nếu capacity của slice nhỏ hơn 1024, Go sẽ tự động gấp đôi capacity của slice mỗi khi cần thêm phần tử. Sau đó, tăng dần 25% khi cần thiết.

- **Giảm thiểu số lần tăng capacity**: Để giảm thiểu thời gian cần thiết khi Go phải cấp phát và sao chép dữ liệu, bạn có thể khai báo slice với một capacity lớn hơn nếu bạn dự đoán số lượng phần tử trong tương lai.

- **Slice từ một slice khác**: Khi bạn tạo một slice từ một slice khác, hai slice này sẽ chia sẻ cùng một vùng nhớ. Ví dụ:

  ```go
  x := []int{1, 2, 3, 4}
  y := x[1:3]  // y = [2, 3]
  ```

- **`copy` function**: Nếu bạn muốn sao chép dữ liệu từ một slice này sang một slice khác mà không thay đổi vùng nhớ của chúng, bạn có thể sử dụng hàm `copy`. Ví dụ:

  ```go
  x := []int{1, 2, 3, 4}
  y := make([]int, 4)
  copy(y, x)
  ```

#### Các lưu ý khi làm việc với Slice

- **Tránh việc sử dụng `append` với subslice**: Khi bạn append vào một subslice (slice được tạo từ slice khác), có thể xảy ra tình huống bạn vô tình ghi đè (overwrite) dữ liệu. Để tránh điều này, bạn có thể sử dụng full slice expression hoặc `copy` thay vì `append`.

- **Sử dụng full slice expression**: Full slice expression cho phép bạn chỉ định rõ ràng phần length và capacity của slice khi tạo một subslice. Ví dụ:

  ```go
  y := x[:2:2]
  z := x[2:4:4]
  ```

  Điều này giúp giới hạn capacity của subslice bằng length.

### Tóm lại

- Sử dụng **mảng** khi bạn biết chính xác số lượng phần tử và kích thước của mảng không thay đổi.
- Sử dụng **slice** khi bạn cần làm việc với một danh sách có thể thay đổi kích thước hoặc không biết chính xác số lượng phần tử.
- Đảm bảo khai báo slice với capacity hợp lý để tránh việc phải cấp phát lại bộ nhớ quá nhiều lần.

---

## Strings và Runes và Bytes

- 1 string = 1 chuỗi các byte

## Maps

map[keyType]valueType

- `ok` là 1 biến dùng để chỉ sữ khác biệt giữa 1 key được gắn với zero value và 1 key không có trong map
- Thường khi gán biến từ map, ta sẽ khai báo 2 biến: đầu tiên là giá trị của key, thứ hai là giá trị bool, thường đặt là `ok`. Nếu `ok` true thì key có trong map, và ngược lại

- Xoá key-value bằng delete(map, key)

## Struct

- Maps không phải 1 cách lý tưởng để truyền data từ function qua function. Khi bạn có dữ liệu liên quan đến nhau mà muốn nhóm lại, nên dùng struct

- Go không có class để kế thừa nhưng không có nghĩa Go không có OOP, chỉ là nó hơi khác đi thôi

- Có thể so sánh các thành phần trong struct tuỳ vào type của nó

## for, 4 cách

- Hoàn chỉnh, C-style for
  for i := 0; i < 10; i++ {}
- 1 điều kiện for
  i := 1
  for i < 100 {}
- for vô hạn
  for {}
- for-range
  for k, v := range x {}

## Functions

- Go không có tham số named và optional input
- Có thể trả nhiều output
- Bắt buộc phải trả về đủ số lượng output
- Nếu hàm trả về giá trị, không bao giờ dùng blank return

### Variadic Input Parameters and Slices

- Tham số này phải ở vị trí cuối cùng trong danh sách tham số hàm. Biến được tạo trong hàm là slice có type là type truyền vào.

- Ví dụ:

```bash
func addTo(base int, vals ...int) []int {
  out := make([]int, 0, len(vals))
  // Do something
  return out
}
```

### Anonymous Functions

- Khai báo bên trong 1 func khác và gán chúng với biến. Tuy nhiên không cần phải gán chúng với 1 biến

- Ví dụ:

```bash
func main() {
  for i := 0; i < 5; i++ {
  func }
}
```

### Closures

- Hàm được khai báo trong hàm gọi là closure. Nghĩa là hàm được khai báo bên trong (inner function) này có thể truy cập và thay đổi giá trị biến được khai báo ở hàm bên nogaif (outer function)
- Mục đích là giới hạn scope của hàm. Hạn chế số lượng khai báo ở package level, dùng inner function sẽ ẩn đi được lời gọi hàm

### Defer

- Keyword dùng để đảm bảo 1 hàm hoặc statement được thực thi ở cuối cùng của function, trước khi nó exit
- Program thường có các tài nguyên tạm, ví dụ như file hoặc kết nối mạng cần được dọn dẹp. Trong Go, việc cleanup này được attach tới function với keyword `defer`
- Defer chạy theo thứ tự last-in-first-out
- Code trong defer closure chạy sau lệnh `return`. Các biến truyền vào 1 defer closure sẽ không được evaluate cho tới khi closure chạy

---

## Pointers - Con trỏ

- Khi nào nên sử dụng con trỏ
- Cách bộ nhớ được cấp phát
- Cách sử dụng con trỏ và giá trị

### A Quick Pointer Primer

- Con trỏ đơn thuần là 1 biến lưu giữ địa chỉ trong bộ nhớ nơi mà 1 giá trị được lưu tại đó
- Mọi biến được lưu vào 1 hoặc nhiều vùng nhớ liên tiếp, được gọi là địa chỉ (address)
- Zero value của con trỏ là `nil`
- Trong Go không có arithmetic pointer
- Ký tự `&` là toán tử địa chỉ (address operation). Nó dẫn đến 1 value type và trả về địa chỉ của vùng nhớ nơi giá trị được lưu
- Ký tự `*` là toán tử gián tiếp (indirection operator). Nó dẫn đến 1 biến kiểu con trỏ (pointer type) và trả về giá trị được trỏ tới -> Được gọi là `dereferencing`
- Trước khi `dereferencing` 1 con trỏ, cần đảm bảo con trỏ ấy không `nil`
- Pointer type được khai báo với `*` trước kiểu dữ liệu. Ví dụ `var x *int`
- Hàm `new` sẽ tạo 1 biến pointer. Tuy nhiên nó ít được sử dụng. Ví dụ:

  ```go
  var x = new(int)
  fmt.Println(x == nil) // prints false
  fmt.Println(*x) // prints 0
  ```

### Dont's Fear the Pointers

- Khác với JS, Java, Python, ở Go có thể sử dụng giữa con trỏ và giá trị cho cả struct lẫn primitives.
- Dùng con trỏ để chỉ ra rằng 1 tham số có thể thay đổi (modify)
- Go sử dụng hệ thống call-by-value, giá trị truyền vào function sẽ được copy
- Nếu kiểu dữ liệu truyền vào function là kiểu non-pointer (primitives, structs, và arrays), hàm được gọi không thể thay đổi giá trị ban đầu
- Nếu truyền con trỏ, function nhận 1 copy của pointer đó, nghĩa là trỏ vào giá trị gốc. Khi ấy function có thể thay đổi dữ liệu mà con trỏ trỏ tới

### Pointers Indicate Mutable Parameters

- Function không thể làm giá trị nil pointer trỏ tới non-nil
- Để thay đổi giá trị được trỏ tới bởi 1 tham số con trỏ, cần phải `dereference` con trỏ

### Pointers Are a Last Resort

- Khi truyền dữ liệu vào 1 function, nên khởi tạo data type trong function và trả về nó, thay vì truyền 1 pointer vào function
- Nên sử dụng con trỏ khi làm việc với kiểu dữ liệu dùng trong concurrency
- Khi làm việc với JSON, bạn sử dụng pointer để phân biệt giữa zero value với missing value

### The Zero Value Versus No Value

- Con trỏ có thể được sử dụng biểu diễn biến chưa được gán giá trị hoặc 1 field của struct

### The Difference Between Maps and Slices

- Map như 1 con trỏ tới 1 struct. Truyền 1 map vào function nghĩa là đang copy 1 pointer. Vì vậy nên tránh sử dụng `map` cho tham số đầu vào hoặc giá trị trả về bởi nó sẽ thay đổi biến ban đầu (original variable)
- Trong khi đó, sử dụng `append` với slice sẽ không thay đổi biến ban đầu do slice có 3 trường:
  - field `int` cho length
  - field `int` cho capacity
  - 1 con trỏ tới 1 vùng nhớ
- Cho nên khi thay đổi length hoặc capacity của slice sẽ chỉ thay đổi biến copy chứ không thay đổi biến ban đầu (original variable)

### Slices as Buffers

- Sử dụng slice như buffer để giảm công việc của GC (garbage collector) bằng việc giảm cấp phát bộ nhớ

### Reducing the Garbage Collector's Workload

- Cấp phát data cho stack càng nhiều càng tốt -> giảm workload của GC. Tuy nhiên việc tăng stack quá nhiều cũng ảnh hưởng đến hiệu năng

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
