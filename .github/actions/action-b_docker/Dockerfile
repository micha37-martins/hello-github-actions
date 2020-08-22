FROM golang:1.14.6-alpine3.12 as golang-builder

RUN echo PWD:
RUN pwd
RUN echo $HOME
RUN ls -a $HOME
RUN echo $GITHUB_REPOSITORY
RUN echo $GITHUB_WORKSPACE
RUN ls $GITHUB_WORKSPACE/src

# use --no-cache to not cache the index locally
# RUN apk --no-cache add git build-base alsa-lib-dev alsa-utils alsa-utils-doc alsa-lib alsaconf

# go.sum & go.mod only exist if gocode have dependencies
# ADD $GITHUB_WORKSPACE/go.mod .
# ADD $GITHUB_WORKSPACE/go.sum .
COPY Makefile .

RUN make test

COPY cmd/ cmd/
COPY internal/ internal/
COPY soundfiles/ soundfiles/

RUN make test
