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
