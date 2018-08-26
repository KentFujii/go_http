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

ENV OPENSSL_CNF /etc/ssl/openssl.cnf
ENV OPENSSL_SUBJ "/C=JP/ST=Tokyo/L=Minato-ku/O=KentFujii/OU=Development/CN=localhost/"
RUN mkdir /certs/
ADD openssl.cnf "${OPENSSL_CNF}"

RUN openssl genrsa -out /certs/ca.key 2048
RUN openssl req -new -sha256 -key /certs/ca.key -out /certs/ca.csr -config "${OPENSSL_CNF}" -subj "${OPENSSL_SUBJ}"
RUN openssl x509 -in /certs/ca.csr -days 365 -req -signkey /certs/ca.key -sha256 -out /certs/ca.crt -extfile "${OPENSSL_CNF}" -extensions CA

RUN openssl rsa -in /certs/ca.key -text
RUN openssl req -in /certs/ca.csr -text
RUN openssl x509 -in /certs/ca.crt -text

RUN openssl genrsa -out /certs/server.key 2048
RUN openssl req -new -nodes -sha256 -key /certs/server.key -out /certs/server.csr -config "${OPENSSL_CNF}" -subj "${OPENSSL_SUBJ}"
RUN openssl x509 -req -days 365 -in /certs/server.csr -sha256 -out /certs/server.crt -CA /certs/ca.crt -CAkey /certs/ca.key -CAcreateserial -extfile "${OPENSSL_CNF}" -extensions Server
