FROM golang:1.21-alpine AS builder

LABEL author="Mystic <215104920@qq.com>"

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
