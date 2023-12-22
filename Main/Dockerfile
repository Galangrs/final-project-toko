FROM golang

WORKDIR /Bot-Discord

COPY . .

RUN go mod download


CMD ["go", "run", "main.go"]
