# Specify Go version
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /gp/src/app

# Copy the project files into the container
COPY . .

# Build the Go application
RUN go build -o main ./cmd

# Expose the necessary port
EXPOSE 8080

# Run the executable
CMD ["./main"]
