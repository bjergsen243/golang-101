# Channels

## Default Channel

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

## Buffered Channels

- Theo mặc định thì channels đều unbuffered, nghĩa là chúng sẽ chỉ nhận gửi (chan <-) nếu có 1 nhận tương ứng (<- chan) sẵn sàng để nhận giá trị được gửi. Buffered channels chấp nhận 1 số lượng giá trị được giới hạn mà không cần receiver tương ứng
- Ví dụ:

  ```go
  messages := make(chan string, 2) // channel này được buffer để nhận tới 2 giá trị
  ```

## Channel Synchronization

- Ta sử dụng channel để đồng bộ thực thi trên toàn goroutines.
- Đây là ví dụ sử dụng blocking receive để chờ 1 goroutine hoàn thành
- Khi chờ nhiều goroutine hoàn thành, nên sử dụng WaitGroup
