FROM golang:1.21-alpine
WORKDIR /app
COPY menu_service.go .
RUN go mod init menu_service
RUN go build -o menu_service
EXPOSE 8081
CMD ["./menu_service"]