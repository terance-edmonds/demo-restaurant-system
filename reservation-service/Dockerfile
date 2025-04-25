FROM golang:1.21-alpine
WORKDIR /app
COPY main.go .
RUN go mod init reservation_service
RUN go build -o reservation_service
EXPOSE 8085
CMD ["./reservation_service"]