FROM golang:1.23.5-alpine3.21 AS builder
WORKDIR /app
COPY go.mod go.sum main.go ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o otus-go-app .

FROM alpine:3.21.2
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/otus-go-app ./
CMD ["/root/otus-go-app"]