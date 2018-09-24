FROM golang:1.9-stretch
RUN apt-get update && apt-get install --no-install-recommends -y \
    gcc \
    build-essential \
    ca-certificates \
    openssl \
    curl \
    git-core

ENV GOROOT /usr/local/go
ENV GOBIN /usr/local/go/bin
ENV GOPATH /go
ENV PATH $PATH:/usr/local/go/bin
RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/derekparker/delve/cmd/dlv
WORKDIR /go/src/app
