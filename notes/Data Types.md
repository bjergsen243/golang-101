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
