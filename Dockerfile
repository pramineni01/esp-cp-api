FROM golang:latest as builder

WORKDIR /esp-cp-api
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -i -installsuffix cgo -o esp-cp-api cmd/cp-api/server.go

FROM alpine:latest

WORKDIR /opt/esp-cp-api/
COPY --from=builder /esp-cp-api/esp-cp-api esp-cp-api
COPY ./cmd/cp-api/config ./config

ENTRYPOINT ["./esp-cp-api"]