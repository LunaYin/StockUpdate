# StockUpdate
### compile proto files
```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=./frontend --proto_path=. service.proto
```
### local run:
```shell
source env
go run cmd/stockupdate.go
```
request:
```shell
curl -s http://localhost:9000/stocks/icn
curl -X POST -s -H "Content-Type: application/json" -d '{"warehouseStock": {"WarehouseUid": "test1", "Quantity": 500}}' http://localhost:9000/stocks/icn/aggregate
```