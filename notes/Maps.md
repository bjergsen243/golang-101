## Maps

map[keyType]valueType

- `ok` là 1 biến dùng để chỉ sữ khác biệt giữa 1 key được gắn với zero value và 1 key không có trong map
- Thường khi gán biến từ map, ta sẽ khai báo 2 biến: đầu tiên là giá trị của key, thứ hai là giá trị bool, thường đặt là `ok`. Nếu `ok` true thì key có trong map, và ngược lại

- Xoá key-value bằng delete(map, key)