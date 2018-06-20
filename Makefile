gen:
	docker-compose run grpc protoc --proto_path=. --go_out=plugins=grpc-aerospike-kvs:./ grpc-aerospike-kvs.proto

up:
	docker-compose up -d
