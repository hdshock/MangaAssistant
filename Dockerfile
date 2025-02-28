# Use the official Golang image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o manga-assistant .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./manga-assistant"] 