#### 模擬搶單實驗

目錄結構
```
|GrabOrders
|+-- client
|   +-- client.go
|+-- pool
|   +-- repo
|       +-- repo.go
|   +-- pool.go
|+-- database
|   +-- db.go
|   +-- order.schema
|+-- go.mod
|+-- go.sum
```

前提:
1. 透過 docker 執行 mysql 容器，建立 sample_db

實驗步驟:
1. 到 pool/ 底下 go run pool.go gen 100000 產生 10000 筆新訂單
2. 到 client/ 底下 go run client.go 觀察 10 個客戶搶奪 10000 筆訂單

核心參考: \
[MySQL row level lock](https://github.com/crossgate10/Today-I-Learned/blob/master/20191124005900.md)
