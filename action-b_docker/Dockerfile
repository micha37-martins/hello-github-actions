FROM golang:1.14.6-alpine3.12 as golang-builder

RUN echo PWD:
RUN pwd
RUN echo GITHUBWORKSPACE:
RUN ls -a ../$GITHUB_WORKSPACE

# use --no-cache to not cache the index locally
# RUN apk --no-cache add git build-base alsa-lib-dev alsa-utils alsa-utils-doc alsa-lib alsaconf

# go.sum & go.mod only exist if gocode have dependencies
# ADD $GITHUB_WORKSPACE/go.mod .
# ADD $GITHUB_WORKSPACE/go.sum .
ADD Makefile .

RUN make test

ADD $GITHUB_WORKSPACE/cmd/ cmd/
ADD $GITHUB_WORKSPACE/internal/ internal/
ADD $GITHUB_WORKSPACE/soundfiles/ soundfiles/

RUN make test
