# URL Shortener (Go + Cloud Native)

## Run locally
go run cmd/server/main.go

## Docker
docker build -t url-shortener .
docker run -p 8080:8080 url-shortener

## Endpoints
- /shorten?url=...
- /r/{code}
- /metrics