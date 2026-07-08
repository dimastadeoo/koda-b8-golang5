FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o program-auth main.go


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/program-auth /app

CMD ["/app/program-auth"]