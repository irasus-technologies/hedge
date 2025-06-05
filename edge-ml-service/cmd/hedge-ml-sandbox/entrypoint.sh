#!/bin/sh


# Ensure slirp4netns is used for networking, not working right now, so used privileged=true etc
#export DOCKERD_ROOTLESS_ROOTLESSKIT_NET=slirp4netns

# Start Docker daemon in background
dockerd-entrypoint.sh &

# Wait until the Docker daemon is up (simple check)
while ! docker info > /dev/null 2>&1; do
  echo "Waiting for dockerd to start..."
  sleep 1
done

echo "Docker daemon is up."


# Run your main binary with arguments passed from CMD or docker-compose `command:`
exec /hedge-ml-sandbox "$@"
