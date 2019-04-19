build:
	protoc -I. --go_out=plugins=micro:. proto/user/user.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t shippy-service-user .
run:
	docker run -p 8082:8080 -e MICRO_SERVER_ADDRESS=:8080 shippy-service-user