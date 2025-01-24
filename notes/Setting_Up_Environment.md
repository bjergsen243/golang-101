# Cài đặt môi trường cho Go

- TLDR:
  - Cài đặt thì xem [tại đây](https://go.dev/dl/)
  - go run: chạy ứng dụng, xoá file binary sau khi chạy
  - go build: build ứng dụng và tạo file binary
  - gofmt: format lại code
  - go lint: yêu cầu code tuân thủ các quy tắc
  - go vet: tìm lỗi trong code
  - golangci-lint: kết hợp của `go lint` và `go vet`
  - [Go's blog](https://go.dev/blog/)

## Cài đặt Go Tools

- Tải Go phù hợp với hệ điều hành của bạn [tại đây](https://go.dev/dl/)
  - Với Mac có thể cài đặt qua Homebrew: `brew install go`
  - Với Windows có thể cài đặt qua Chocolatey: `choco install golang`
- Kiểm tra Go đã được cài đặt hay chưa:

  ```bash
  go version
  ```

## Go Workspace

- Chú ý: Cần thêm GOPATH và binary của Go vào PATH qua lệnh sau:

  ```bash
  export GOPATH=$HOME/go
  export PATH=$PATH:$GOPATH/bin
  ```

- Workspace đặt ở thư mục `$HOME/go`.
- Source code cho các tools ở `$HOME/go/src`.
- Mã nhị phân được biên dịch ở `$HOME/go/bin`.

## Câu lệnh `go`

Các câu lệnh `go` hay được sử dụng

### go run

- Lệnh `go run file_name.go` sẽ biên dịch code thành nhị phân. Tuy nhiên nhị phân này được lưu ở thư mục tạm thời, sau khi thực thi xong sẽ bị xoá.
- Sử dụng lệnh này khi muốn chạy code ngay lập tức.

### go build

- Lệnh `go build hello.go` hoặc `go build -o hello_world hello.go` sẽ tạo ra 1 file nhị phân trong thư mục hiện tại.
- Sử dụng lệnh này để tạo file nhị phân mà sẽ đưa cho người khác sử dụng.

### Cài đặt công cụ bên thứ 3

- Khi bạn muốn cài đặt các thư viện không có sẵn, sử dụng lệnh `go install link_to_library` để cài đặt thư viện đó.
- Nó sẽ tải thư viện, các thành phần của nó và cài đặt binary vào thư mục `$GOPATH/bin`.

### Format Code

- Đơn giản với lệnh `go fmt ./...` hoặc `gofmt` sẽ format lại source code. Tham khảo thêm [tại đây](https://pkg.go.dev/cmd/gofmt).
- Có thể tham khảo [goimports](https://pkg.go.dev/golang.org/x/tools@v0.29.0/cmd/goimports).
- Chú ý: Luôn chạy `go fmt` trước khi biên dịch code.

### Linting và Vetting

- Ngoài `go fmt` giúp ta format code thì có các công cụ giúp việc viết code sẽ tuân thủ theo các quy tắc nhất định là `golint`.
- Cài đặt `golint` qua lệnh:

  ```bash
  go install golang.org/x/lint/golint@latest
  ```

- Và chạy nó với:

  ```bash
  golint ./...
  ```

- Có thêm lệnh `go vet` để xác định các lỗi như truyền sai tham số hoặc gán tới 1 biến không bao giờ được sử dụng (lỗi này sẽ được xác định lúc compile time rồi)
- Chạy nó với:

  ```bash
  go vet ./...
  ```

- Bên cạnh đó cũng có 1 tool khác là [golangci-lint](https://github.com/golangci/golangci-lint). Nó chạy 10 linter khác nhau, ta có thể sửa lên 50 cái.

- Tuy nhiên, nên sử dụng `go vet` như 1 phần yêu cầu khi build, còn `go lint` sẽ là 1 phần trong quá trình review code (do `golint` có thể nhiều trường hợp false positives và false negatives, nên không thể yêu cầu fix hết chúng được). Khi đã quen với gợi ý của chúng rồi thì có thể thử dùng `golangci-lint`.

## IDE

- Tuỳ vào sự thuận tiện khi code của bạn, có thể sử dụng Visual Studio Code (miễn phí), GoLand (dùng thử), [The Go Playground](https://go.dev/play/) (dùng online trên web).

## Hãy nhớ cập nhật phiên bản mới

- Có 1 điều mà bản thân tôi thấy khi sử dụng Go là nên chú ý tới việc cập nhật phiên bản mới cho nó, bởi vì sẽ có nhiều tính năng mới khá hay ho. Ví dụ như gần đây đã cập nhật Generics (1.18)

- Các tính năng của phiên bản mới sẽ được cập nhật tại [blog của Go](https://go.dev/blog/).
