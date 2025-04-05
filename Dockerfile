FROM golang:1.24.1-alpine AS builder

COPY . /github.com/nikita-reshetnyak/chat-server/source
WORKDIR /github.com/nikita-reshetnyak/chat-server/source

RUN go mod download
RUN go build -o ./bin/chat-server cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/nikita-reshetnyak/chat-server/source/bin/chat-server .

CMD ["./chat-server"]