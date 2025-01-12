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
