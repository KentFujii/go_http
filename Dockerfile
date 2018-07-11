FROM ubuntu:16.04
RUN apt-get update && apt-get install --no-install-recommends -y \
    gcc \
    build-essential \
    ca-certificates \
    curl \
    mercurial \
    git-core
RUN curl -s https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz| tar -v -C /usr/local -xz
ENV GOROOT /usr/local/go
ENV GOBIN /usr/local/go/bin
ENV PATH $PATH:/usr/local/go/bin
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN go get -u github.com/derekparker/delve/cmd/dlv

ENV GOPATH /go
WORKDIR /go/src/app
