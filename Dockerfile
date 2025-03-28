FROM golang:1.21 AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o backend .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/backend .
COPY --from=builder /app/.env .  

EXPOSE 8080
CMD ["./backend"]