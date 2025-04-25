FROM golang:1.21-alpine
WORKDIR /app
COPY payment_service.go .
RUN go mod init payment_service
RUN go build -o payment_service
EXPOSE 8084
CMD ["./payment_service"]