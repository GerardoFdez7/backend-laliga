FROM golang:1.24.1-alpine
RUN apk add --no-cache postgresql-client dos2unix
RUN go install github.com/air-verse/air@latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN dos2unix /app/start.sh && \
    chmod +x /app/start.sh
EXPOSE 8080
CMD ["sh", "/app/start.sh"]