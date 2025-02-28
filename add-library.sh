#!/bin/bash

# Define the Windows path to the manga library
WINDOWS_PATH="C:\\Users\\jason.mesisco\\Downloads\\Manga"

# Escape special characters in the path
ESCAPED_PATH=$(echo "$WINDOWS_PATH" | sed 's/\\/\\\\/g' | sed 's/!/\\!/g' | sed 's/@/\\@/g')

# Convert the Windows path to a Docker-friendly format
DOCKER_PATH=$(echo "$ESCAPED_PATH" | sed 's/\\\\/\//g' | sed 's/^C:/\/c/')

# Create the Docker volume with the bind mount
docker volume create manga_library --driver local --opt type=none --opt device="$DOCKER_PATH" --opt o=bind

# Update the docker-compose.yml file
if ! grep -q "manga_library" docker-compose.yml; then
    echo "Adding manga_library volume to docker-compose.yml..."
    sed -i '/volumes:/a \  manga_library:' docker-compose.yml
fi

echo "Volume created and docker-compose.yml updated successfully."
