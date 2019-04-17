build:
	protoc -I. --go_out=plugins=micro:. proto/vessel/vessel.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t shippy-service-vessel .
run:
	docker run -p 8081:8081 -e MICRO_SERVER_ADDRESS=:8081 shippy-service-vessel