# Use an official Golang runtime as the base image
FROM golang:1.16

# Set the working directory in the container to /app
WORKDIR /app

# Copy go.mod and go.sum to the working directory
COPY go.mod go.sum ./

# Download all the dependencies
RUN go mod download

# Copy the rest of the application code to the working directory
COPY . .

# Build the application
RUN go build -o main .

# Expose port 4700 for the application
EXPOSE 4700

# Start the application
CMD ["./main"]