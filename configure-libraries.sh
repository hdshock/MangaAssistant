#!/bin/bash

# This script mounts manga library paths directly into the Docker container
# and then starts the Docker containers.

# Default libraries file
LIBRARIES_FILE="libraries.txt"

# Check if the libraries file exists
if [ ! -f "$LIBRARIES_FILE" ]; then
  echo "Error: $LIBRARIES_FILE not found."
  echo "Please create a file named $LIBRARIES_FILE and add your library paths (one per line)."
  echo "Example:"
  echo "C:\Manga\Library1"
  echo "\\server\share\Library2"
  echo "/home/user/manga"
  exit 1
fi

# Function to convert Windows paths to Docker-friendly format
convert_path() {
  local path=$1
  if [[ "$path" == *\\* ]]; then
    path=$(echo "$path" | sed 's/\\/\//g' | sed 's/^\([A-Za-z]\):/\/\1/')
    echo "$path"
  else
    echo "$path"
  fi
}

# Read paths from the libraries file (one path per line)
LIBRARY_MOUNTS=()
while IFS= read -r path; do
  if [ -n "$path" ]; then
    docker_path=$(convert_path "$path")
    LIBRARY_MOUNTS+=("--mount type=bind,source=$docker_path,target=/manga-library/$(basename "$docker_path")")
    echo "Added library: $docker_path"
  fi
done < "$LIBRARIES_FILE"

echo "Libraries configured successfully. Starting Docker containers..."

# Start the Docker containers with the mounted libraries
docker-compose down
docker-compose up --build ${LIBRARY_MOUNTS[@]} 