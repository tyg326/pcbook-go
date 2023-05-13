# windows下make命令使用mingw64中的mingw32-make.exe
#rm pb/*.go mingw32-make默认使用的是cmd命令, 即使用del代替rm,同时路径还要按照\分割#
# del pb/*.go 提示无效开关 - "*.go" 解决: del pb\*.go
gen:
# 第一种
	protoc --proto_path=proto --go_out=./pb --go-grpc_out=./pb proto/*.proto
# 第二种 指定makefile目录为路径为工作路径 import proto时要加proto路径
# protoc --proto_path=. \
# --go_out=. --go_opt=paths=import \
# --go-grpc_out=. --go-grpc_opt=paths=import \
# proto/*.proto

clean:
	del pb\*.go

# run:
# 	go run main.go

server:
	go run cmd/server/main.go --port 18080

client:
	go run cmd/client/main.go --address 0.0.0.0:18080

test:
	go test -cover -race ./...
	
