build:
	protoc -I. --go_out=plugins=micro:. proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-cli .
run:
	docker run -e MICRO_REGISTRY=mdns consignment-cli