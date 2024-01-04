# Start from a base Go image
FROM golang:1.17

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod .
COPY go.sum .

# Fetch dependencies using go mod.
RUN go mod download
RUN go mod verify

COPY . .

# Install the migrate tool
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2

# Set Postgres URL
ARG POSTGRESQL_URL=postgres://postgres:password@172.17.0.1:5432/log?sslmode=disable

# Run the migrations
RUN migrate -database ${POSTGRESQL_URL} -path migrations up

# Build the Go app
RUN go build -o /go/bin/main ./cmd/server/.

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/go/bin/main"]