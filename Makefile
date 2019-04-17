build:
	protoc -I. --go_out=plugins=micro:. vessel-service/proto/vessel/vessel.proto
	protoc -I. --go_out=plugins=micro:. consignment-service/proto/consignment/consignment.proto