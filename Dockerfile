# Use the official Go image as the base
FROM golang:1.17

# Set the working directory
WORKDIR /app

# Copy your Go source code into the container
COPY . .

# Build the Go binary
RUN go build -o myapp

# Expose the port your application listens on
EXPOSE 80

# Command to run your application
CMD ["./myapp"]
