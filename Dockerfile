FROM golang:1.24-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o /out/commitment ./cmd/server
FROM alpine:3.20
WORKDIR /app
COPY --from=build /out/commitment .
EXPOSE 8080
ENTRYPOINT ["/app/commitment"]
