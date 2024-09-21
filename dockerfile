# stage 1
FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/server

EXPOSE 8080

# stage 2
FROM scratch

COPY --from=builder /app/main /
COPY app.env .

ENTRYPOINT [ "/main" ]