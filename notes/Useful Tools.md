# Công cụ hữu ích

## Lint

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

## Kiểm tra lỗi

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

## Makefile

### Thành phần

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

### Lợi ích của Makefile

- Tự động hóa các lệnh phức tạp.
- Đảm bảo tính nhất quán khi chạy lệnh trên các máy khác nhau.
- Giúp quản lý quy trình build và kiểm tra mã nguồn dễ dàng hơn.
