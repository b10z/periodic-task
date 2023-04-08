FROM golang:1.20.2-alpine3.16 as builder

RUN apk update
RUN apk add --no-cache git openssh

WORKDIR /app

COPY ../../go.mod go.sum ./

RUN go mod download

COPY ../.. .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go

FROM golang:1.20.2-alpine3.16

RUN apk --no-cache add ca-certificates

WORKDIR /app/
COPY --from=builder /app/cmd/main .

EXPOSE 8000
EXPOSE 9000

CMD ["./cmd/main"]
