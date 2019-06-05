# build stage
FROM golang:1.12 as builder

ENV GO111MODULE=on

WORKDIR /go/src/github.com/asxcandrew/galas

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/galas cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/migrate migrations/*.go

# final stage
FROM alpine:3.8
COPY --from=builder /go/src/github.com/asxcandrew/galas/build/galas /app/
COPY --from=builder /go/src/github.com/asxcandrew/galas/build/migrate /app/

ENTRYPOINT ["/app/galas"]
