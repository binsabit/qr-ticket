FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app
RUN apk --no-cache add tzdata

COPY internal .

RUN go build -v  -ldflags="-w -s" -o main ./cmd/main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Almaty
COPY --from=builder /app .

CMD ["sh", "-c", "/app/main"]