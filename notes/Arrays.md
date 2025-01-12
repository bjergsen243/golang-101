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
