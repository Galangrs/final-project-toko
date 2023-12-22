FROM golang

WORKDIR /Main

COPY . .

RUN go mod download && go mod tidy

CMD ["go", "run", "main.go"]
