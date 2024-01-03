LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc


generate:
	make generate-note-api

generate-note-api:
	mkdir -p pkg/note_v1
	protoc --proto_path api/note_v1 \
	--go_out=pkg/note_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/note_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/note_v1/note.proto

generate-note-api-me://мой
	//$ protoc -I api/note_v1 api/note_v1/note.proto 
	--go_out=./pkg/note_v1 --go_opt=paths=source_relative 
	--go-grpc_out=./pkg/note_v1/ --go-grpc_opt=paths=source_relative

go build:
	go build -o ./bin/ cmd/grpc_server/main.go
for bash:
	GOOS=linux GOARCH=amd64 go build -o service_linux cmd/grpc_server/main.go //мой

	ssh root@185.91.52.118
	6uGuus0FtHYV

	scp service_linux root@185.91.52.118: #копируем в облако

	sudo chmod 755 service_linux #доступ в линксе от админа

docker-build-and-push:
	docker buildx build --no-cache --platform linux/amd64 -t cr.selcloud.ru/amani/test-server:v0.0.1 .
	docker login -u token -p CRgAAAAAttY9irS1XHQYOzSTkesCGQPsUqpBYlUs cr.selcloud.ru/amani
	docker push cr.selcloud.ru/amani/test-server:v0.0.1