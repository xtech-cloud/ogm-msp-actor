.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/actor/healthy.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/actor/shared.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/actor/domain.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/actor/device.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/actor/sync.proto

