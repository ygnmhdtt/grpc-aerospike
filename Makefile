gen:
	docker-compose run grpc protoc --go_out=plugins=grpc:. grpc-aerospike-kvs.proto

up:
	docker-compose up -d

build:
	docker-compose build

stop:
	docker-compose stop
