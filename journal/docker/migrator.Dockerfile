FROM golang:1.26.1

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY migrations /migrations
RUN GOOS=linux go build -o /migrator ./cmd/migrator/main.go

ENTRYPOINT ["/migrator"]
