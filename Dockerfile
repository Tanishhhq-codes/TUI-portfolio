FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o portfolio .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/portfolio .

# SSH host key will be generated here on first run
RUN mkdir -p .ssh

EXPOSE 23234

CMD ["./portfolio"]
