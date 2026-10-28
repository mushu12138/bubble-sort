# bubble-sort

Go 实现的冒泡排序，附带完整单元测试。

## 运行

```bash
go run bubble.go
```

## 测试

```bash
go test -v ./...
```

## 说明

- `BubbleSort` 原地排序，含提前退出优化（一轮无交换即视为有序）
- 测试覆盖：普通数组、已有序、逆序、重复元素、负数、单元素、空切片
