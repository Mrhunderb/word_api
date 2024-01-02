FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]

