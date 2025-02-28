#!/bin/bash

# Create the Go module files
if [ ! -f go.mod ]; then
    echo "Initializing Go module..."
    go mod init manga-assistant
    echo "package main\n\nfunc main() {}\n" > main.go
    go mod tidy
else
    echo "Go module already exists."
fi

# Create the frontend directory and initialize a React app
if [ ! -d frontend ]; then
    echo "Setting up React frontend..."
    mkdir -p frontend
    echo '{
      "name": "manga-assistant-frontend",
      "version": "1.0.0",
      "private": true,
      "dependencies": {
        "react": "^18.2.0",
        "react-dom": "^18.2.0",
        "react-router-dom": "^6.18.0",
        "axios": "^1.6.2",
        "tailwindcss": "^3.3.3"
      },
      "scripts": {
        "start": "react-scripts start",
        "build": "react-scripts build",
        "test": "react-scripts test",
        "eject": "react-scripts eject"
      },
      "devDependencies": {
        "autoprefixer": "^10.4.16",
        "postcss": "^8.4.31",
        "react-scripts": "^5.0.1"
      }
    }' > frontend/package.json
    echo '{}' > frontend/package-lock.json
else
    echo "Frontend directory already exists."
fi

# Build and run the Docker Compose setup
echo "Building and starting Docker Compose..."
docker-compose down
docker-compose up --build 