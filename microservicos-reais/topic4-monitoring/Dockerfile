FROM golang:1.23 as builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM scratch
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8081
CMD ["./main"]  