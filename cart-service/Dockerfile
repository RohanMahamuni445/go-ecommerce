FROM golang:1.23-alpine AS firstphase

WORKDIR /app

COPY go.mod go.sum .

RUN go mod download

COPY . .

RUN  go mod tidy

RUN go build -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=firstphase /app/main .

EXPOSE 8083

CMD ["./main"]

