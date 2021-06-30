# StockUpdate
local run:
```shell
source env
go run cmd/stockupdate.go
```
request:
```shell
curl -s http://localhost:9000/stocks/icn
curl -X POST -s -H "Content-Type: application/json" -d '{"warehouseStock": {"WarehouseUid": "test1", "Quantity": 500}}' http://localhost:9000/stocks/icn/aggregate
```