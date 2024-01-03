FROM golang:1.20.3-alpine AS builder

COPY . /github.com/Pavel910862/microservices_course/week_1/grpc/source/
WORKDIR /github.com/Pavel910862/microservices_course/week_1/grpc/source/

RUN go mod download
RUN go build -o ./bin/crud_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/Pavel910862/microservices_course/week_1/grpc/source/bin/crud_server .

CMD ["./crud_server"]