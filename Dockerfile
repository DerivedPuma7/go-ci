FROM golang:1.21.5

WORKDIR /app

RUN go mod init github.com/derivedpuma7/ci

COPY . .

RUN go build -o math

CMD ["./math"]