FROM ubuntu:16.04
RUN apt-get update && apt-get install --no-install-recommends -y \
    gcc \
    build-essential \
    ca-certificates \
    openssl \
    curl \
    git-core
RUN curl -s https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz| tar -v -C /usr/local -xz
ENV GOROOT /usr/local/go
ENV GOBIN /usr/local/go/bin
ENV PATH $PATH:/usr/local/go/bin
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN go get -u github.com/derekparker/delve/cmd/dlv

ENV GOPATH /go
WORKDIR /go/src/app

ENV OPENSSL_CNF /etc/ssl/openssl.cnf
ENV OPENSSL_SUBJ "/C=JP/ST=Tokyo/L=Minato-ku/O=KentFujii/OU=Development/CN=GoHTTP/"
ADD openssl.cnf "${OPENSSL_CNF}"
RUN openssl genrsa -out ca.key 2048
RUN openssl req -new -sha256 -key ca.key -out ca.csr -config "${OPENSSL_CNF}" -subj "${OPENSSL_SUBJ}"
RUN openssl x509 -in ca.csr -days 365 -req -signkey ca.key -sha256 -out ca.crt -extfile "${OPENSSL_CNF}" -extensions CA

RUN openssl rsa -in ca.key -text
RUN openssl req -in ca.csr -text
RUN openssl x509 -in ca.crt -text

RUN openssl genrsa -out server.key 2048
RUN openssl req -new -nodes -sha256 -key server.key -out server.csr -config "${OPENSSL_CNF}" -subj "${OPENSSL_SUBJ}"
RUN openssl x509 -req -days 365 -in server.csr -sha256 -out server.crt -CA ca.crt -CAkey ca.key -CAcreateserial -extfile "${OPENSSL_CNF}" -extensions Server
