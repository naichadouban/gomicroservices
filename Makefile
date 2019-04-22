build_linux:
	protoc -I. --go_out=plugins=micro:. vessel-service/proto/vessel/vessel.proto
	protoc -I. --go_out=plugins=micro:. consignment-service/proto/consignment/consignment.proto
	protoc -I. --go_out=plugins=micro:. user-service/proto/user/user.proto

	GOOS=linux GOARCH=amd64 go build  -o ./vessel-service/vessel-service ./vessel-service
	GOOS=linux GOARCH=amd64 go build  -o ./consignment-service/consignment-service ./consignment-service
	GOOS=linux GOARCH=amd64 go build  -o ./user-service/user-service ./user-service
	GOOS=linux GOARCH=amd64 go build  -o ./consignment-cli/consignment-cli ./consignment-cli
	GOOS=linux GOARCH=amd64 go build  -o ./user-cli/user-cli ./user-cli

build_win:
	protoc -I. --go_out=plugins=micro:. vessel-service/proto/vessel/vessel.proto
	protoc -I. --go_out=plugins=micro:. consignment-service/proto/consignment/consignment.proto
	protoc -I. --go_out=plugins=micro:. user-service/proto/user/user.proto
	go build  -o ./vessel-service/vessel-service.exe ./vessel-service
	go build  -o ./consignment-service/consignment-service.exe ./consignment-service
	go build  -o ./user-service/user-service.exe ./user-service
	go build  -o ./consignment-cli/consignment-cli.exe ./consignment-cli
	go build  -o ./user-cli/user-cli.exe ./user-cli
