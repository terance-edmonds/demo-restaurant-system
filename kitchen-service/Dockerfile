FROM golang:1.21-alpine
WORKDIR /app
COPY kitchen_service.go .
RUN go mod init kitchen_service
RUN go build -o kitchen_service
EXPOSE 8083
CMD ["./kitchen_service"]