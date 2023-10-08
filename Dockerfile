FROM golang:1.20

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .

CMD ["go", "run", "main.go"]
