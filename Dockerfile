FROM golang:1.19.3-alpine3.15

# Maintainer Info
LABEL maintainer="Sergey Shchetkin <mrschetkin@gmail.com>"

# Install suuport tools
RUN apk add --no-cache bash

# Move to working directory
WORKDIR /go/src/app

# Copy the code into the container
COPY . .

# Download packages
RUN apk add --no-cache --virtual .ext-deps build-base \
    && go mod download \
    && go mod vendor \
    && apk del .ext-deps