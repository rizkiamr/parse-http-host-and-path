# Stage 1: Build the Go application
FROM golang:1.24.1 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /binary/app

# Stage 2: Create the minimal Distroless image
FROM gcr.io/distroless/static:latest
COPY --from=builder /binary/app /app
CMD ["/app"]
