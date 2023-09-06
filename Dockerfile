# Dockerfile
FROM golang:latest

# コードをコピーするディレクトリに移動
WORKDIR /go/src/app

# コードをコンテナ内にコピー
COPY . .

# Build the application
RUN go build -o main .

# Run the application
CMD ["./main"]