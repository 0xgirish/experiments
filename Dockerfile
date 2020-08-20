FROM golang:1.14-alpine as builder

RUN apk add --no-cache git

WORKDIR /app

ADD ./go.mod ./go.mod
ADD ./go.sum ./go.sum

RUN go mod download

ADD ./main.go ./main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fx-amd64 .

FROM alpine:3.12.0

WORKDIR /root

COPY --from=builder /app/fx-amd64 ./fx-amd64

ENTRYPOINT ["./fx-amd64"]
