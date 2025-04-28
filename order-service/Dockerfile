FROM golang:1.21-alpine
WORKDIR /app
COPY main.go .
RUN go mod init order_service
RUN go build -o order_service
EXPOSE 8082
CMD ["./order_service"]