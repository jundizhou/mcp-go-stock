# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum* ./
COPY ../go.mod ../go.mod
COPY ../go.sum ../go.sum

# Copy source code
COPY . .
COPY .. /go-stock

# Set replace directive for parent module
RUN go mod edit -replace go-stock=/go-stock

# Download dependencies
RUN go mod tidy

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o jd-go-stock .

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Install runtime dependencies (for chromedp if needed)
RUN apk add --no-cache ca-certificates chromium

# Copy binary from builder
COPY --from=builder /app/jd-go-stock .

# Expose port for SSE/HTTP modes
EXPOSE 3000

# Default command: SSE mode
CMD ["./jd-go-stock", "--mode=sse", "--port=3000"]
