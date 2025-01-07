# Golang 101

Go là ngôn ngữ static type. Nghĩa là khi biên dịch thì Go cần biết kiểu dữ liệu của các biến và tham số

-> Nếu dùng generics sẽ phá vỡ nguyên tắc này. Nghĩa là khi tới runtime thì mới check được

## Resources

- GoByExample: https://gobyexample.com/
  https://200lab.io/blog/tag/golang/
- sách Learning Go

## Useful Tools

### Lint

Dùng lint:
go lint -
Cài đặt “go install golang.org/x/lint/golint@latest”

Chạy với lệnh

golint ./...

### Kiểm tra lỗi

chạy lệnh

go vet ./...

hoặc dùng tool golangci-lint sẽ chạy đồng thời golint, go vet: https://golangci-lint.run/usage/configuration/

-> Nên dùng golint và go vet. Chỉ dùng golangci-lint khi team đồng ý

### Makefile

Sử dụng Makefile để tự động các lệnh chạỵ, tránh việc lệnh này chạy ở máy của mình mà không chạy được ở máy người khác
Mỗi 1 lệnh gọi là target
.DEFAULT_GOAL là target sẽ chạy không không chỉ định target nào

Từ trước dấu hai chấm (:) là tên của target. Bất kỳ từ nào phía sau đều là các target cần phải chạy trước khi chạy target này
Lệnh `.PHONY` giúp `make` không nhầm lẫn khi tạo 1 thư mục trong project mà có trùng tên với target

## Kiểu dữ liệu nguyên thuỷ và Khai báo

### Kiểu Built-in

- boolean, integer, float, stringÏ

#### The Zero Value

- Mặc định gán zero value cho biến mà được khai báo nhưng không gán giá trị
- Với kiểu integer, float là 0

#### Literals

- Một `literal` có thể là 1 số, ký tự hoặc string
- Integer:
  - có thể khai báo như sau:
  - 1, 1_234 (1234)
  - 0b cho số nhị phân, 0o cho số octal, 0x cho hexadecimal - tuy nhiên khai báo như này rất gây nhầm lẫn
- Float:
  - Dùng dấu (.) để khai báo giá trị. Ví dụ: 6.03e23 -> 6.023^23
  - Có float32 và float64
- Rune literal:
  - biểu diễn ký tự và bao quanh bởi dấu ngoặc đơn ('')
  - Ví dụ: 'a', các số octal, 8/16-bit hexadecimal, các ký tự như '\n',...

-> Tóm lại thì:

- Dùng hệ cơ số 10 để biểu diễn số và hạn chế việc dùng hexadecimal với rune literal
- Nên dùng dấu ngoặc kép ("") để tạo 1 interpreted string literal (ví dụ "Hello World")
- Nếu cần dùng backslash, dấu ngoặc kép, hoặc xuống dòng trong string, thì sử dụng raw string literal. Được bao quanh bởi dấu (``)

- Boolean:

  - Có 2 giá trị true hoặc false

- Kiểu số học (numeric types)
  - 1 `byte` tương đương với `uint8`
  - Còn có `rune` và `uintptr`

-> Tóm lại thì:

-

#### Khai báo

- Không nên khai báo kiểu := khi

  - khởi tạo 1 biến với zero value, thay vào đó, sử dụng var x int
  - khi gán 1 constant chưa có kiểu dữ liệu hoặc 1 literal cho 1 biến và kiểu dữ liệu mặc định của nó không phải kiểu bạn muốn cho biến, khai báo như sau: varr x byte = 20
  - Bởi vì := cho phép gán cả biến mới và cũ, đôi khi nó tạo 1 biến mới trong khi bạn nghĩ đang dùng biến cũ. Trong tình huống này, khai báo rõ ràng các biến mới với `var` sẽ hợp lý hơn

- Hạn chế khai báo biến bên ngoài function vì chúng làm phức tạp luồng phân tích dữ liệu

- Tất cả các biến local cần được read -> hạn chế khai báo biến thừa
- Các biến const có thể được khai báo mà không cần read bởi vì nó được tính toán ở compile time và không gây ảnh hưởng. Nghĩa là nếu nó không được sử dụng, thì sẽ được bao gồm trong binary được compile

## Array - Mảng

- Mảng ít khi được sử dụng trực tiếp trong Go bởi vì:
  - Go sẽ lấy cả `size` của mảng thành 1 phần trong `type` của mảng. Ví dụ [3]int sẽ khác kiểu so với [4]int.
  - Không thể dùng 1 biến dùng làm size của mảng, vì `type` cần được xử lý ở compile time, chứ không phải runtime
  - Không thể ép kiểu 1 mảng thành mảng có kiểu khác. Không thể viết 1 function mà hoạt động với mảng có size bất kỳ cũng như không thể gán 1 mảng có size khác nhau với mảng khác
- Go chỉ có mảng 1 chiều, tuy nhiên có thể tạo mảng nhiều chiều như sau: var x [2][3]int

-> Tóm lại thì:

- Dùng mảng khi biết chính xác length
- Mảng tồn tại vì nó 1 chỗ lưu trữ cho `slice`

### Slices

- Khai báo:
  - slice literal
  - nil zero value
  - make
- Khi bạn cần 1 số giá trị tuần tự, sử dụng slice
- Slice hữu dụng bởi vì length không phải 1 phần của `type` của slice
- Sự khác biệt với mảng là khi khai báo slice sẽ không chỉ định size của nó. Ví dụ: var x = []int{10, 20, 30}
- Không nên tạo 1 slice có capacity nhỏ hơn length vì nó sẽ panic tại runtime

* Tips: Sử dụng [...] tạo mảng, [] tạo slice

- Không thể so sánh slice với nhau, chỉ so sánh slice được với nil

- nil là identifier biểu diễn cho việc thiếu sót của giá trị với 1 số kiểu dữ liệu

- package reflect có hàm DeepEqual có thể so sánh slice

- Mỗi khi truyền 1 tham số vào 1 function, Go tạo 1 copy của giá trị được truyền vào

* Capacity

- Mỗi slice có 1 capacity, là số lượng vùng nhớ liên tục dự trữ. Nó có thể lớn hơn length. Khi length tới ngưỡng capacity thì không thể append thêm vào slice đó nữa. Khi đó Go runtime sẽ cấp phát 1 slice với capacity lớn hơn và trả về slice mới này -> tốn thời gian do quá trình tạo mới, sao chép từ slice cũ sang slice mới

- Từ Go1.14 là sẽ gấp đôi size của slice nếu capacity nhỏ hơn 1024 và tăng dần 25% sau đó

- Mục tiêu của việc khai báo slice là giảm thiểu số lần slice cần tăng size

- Tạo slice bằng 1 slice rỗng: var x = []int{}

-> Tóm lại thì:

- Nếu sử dụng slice như 1 buffer thì sử dụng nonzero length
- Nếu biết chính xác size bạn muốn, có thể chỉ định length và index vào slice
- Trong các trường hợp khác, sử dụng make với zero length và chỉ định capacity -> Nên dùng cách này do hạn chế việc bị panic

- Nếu tạo 1 slice từ 1 slice khác, 2 slice này sẽ chia sẻ cùng 1 vùng nhớ. Nghĩa là nếu thay đổi 1 slice thì sẽ có tác động tới các slice có chung thành phần thay đổi đó

- Hạn chế sử dụng append với subslice (slice được tạo từ slice khác) hoặc đảm bảo append không gây ra việc overwrite bằng cách sử dụng full slice expression

y := x[:2:2]
z := x[2:4:4]

- Việc khai báo này giúp giới hạn capacity của subslice bằng length

- Hoặc có thể dùng copy(dest, src). Với hàm này thì không cần quan tâm capacity mà length quan trọng hơn
  x := []int{1, 2, 3, 4}
  y := make([]int, 4)
  num := copy(y, x)

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
