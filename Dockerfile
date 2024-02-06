# syntax=docker/dockerfile:1
FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ENV GIN_MODE=release
ENV PORT=8080

VOLUME ./words-in-the-bible-db /app/words-in-the-bible-db 

RUN CGO_ENABLED=0 GOOS=linux go build -o /is-in-bible-br /app/cmd/main.go

EXPOSE 8080

CMD ["/is-in-bible-br"]
