FROM ubuntu:16.04
RUN apt-get update && apt-get install --no-install-recommends -y \
    ca-certificates \
    curl \
    mercurial \
    git-core
RUN curl -s https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz| tar -v -C /usr/local -xz
ENV GOROOT /usr/local/go
ENV GOBIN /usr/local/go/bin
ENV GOPATH /go
ENV PATH $PATH:/usr/local/go/bin
WORKDIR /go/src/app
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
