# build stage
FROM golang as builder

ENV GO111MODULE=on

WORKDIR /go/src/github.com/asxcandrew/galas
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/galas cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/migrate migrations/*.go

# final stage
FROM alpine:3.8
COPY --from=builder /go/src/github.com/asxcandrew/galas/build/galas /app/
COPY --from=builder /go/src/github.com/asxcandrew/galas/build/migrate /app/

ENTRYPOINT ["/app/galas"]
