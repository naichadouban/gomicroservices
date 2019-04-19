build:
	# 编译proto文件
	protoc -I. --go_out=plugins=micro:. proto/consignment/consignment.proto
	# 编译服务代码
	GOOS=linux GOARCH=amd64 go build
	# 创建docker镜像
	docker build -t shippy-service-consignment .
run:
	# 启动docker容器
	docker run -p 8010:8010 -e MICRO_SERVER_ADDRESS=:8010 shippy-service-consignment
