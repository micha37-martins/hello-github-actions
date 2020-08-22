# This Dockerfile adds required dependencies to the alpine image
FROM golang:1.14.6-alpine3.12 as golang-builder

# use --no-cache to not cache the index locally
RUN apk --no-cache add git build-base alsa-lib-dev alsa-utils alsa-utils-doc alsa-lib alsaconf

# Copies your code file from your action repository to the filesystem path `/` of the container
COPY entrypoint.sh /entrypoint.sh

# Code file to execute when the docker container starts up (`entrypoint.sh`)
ENTRYPOINT ["/entrypoint.sh"]
