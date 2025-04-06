# Step 1: Build stage
FROM golang:1.22 AS build

WORKDIR /app

# Copy go.mod and go.sum first (for better layer caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app

# Step 2: Final stage
FROM scratch

COPY --from=build /app/app /app

EXPOSE 8083

ENTRYPOINT ["/app"]

