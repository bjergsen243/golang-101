# Channels

- Channels là pipe để kết nối đồng thời (concurrent) goroutines
- Có thể gửi giá trị vào channel từ 1 goroutine này và nhận giá trị đó tại 1 goroutine khác
- Tạo mới 1 channel bằng keyword `chan` và `make`. Ví dụ: `messages := make(chan string)`
- Gửi giá trị vào 1 channel sử dụng syntax `<-`. Ví dụ

  ```go
  go func() {messages <- "ping"}()
  ```

- Nhận giá trị từ channel cũng sử dụng syntax `<-`. Ví dụ

  ```go
  msg := <-messages
  ```

- Theo mặc định gửi và nhận chặn cho tới khi sender và receiver sẵn sàng. Nó cho phép chúng ta chờ cho tới cuối chương trình để nhận giá trị mà tránh việc có bất kỳ sự đồng bộ (synchronization) nào khác
