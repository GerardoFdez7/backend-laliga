FROM golang:1.24.1-alpine
RUN apk add --no-cache postgresql-client
WORKDIR /app
COPY go.mod go.sum ./
RUN go get -d ./... && go mod tidy
COPY . .
COPY start.sh /app/start.sh
RUN chmod +x /app/start.sh
EXPOSE 8080
CMD ["sh", "/app/start.sh"]