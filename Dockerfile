FROM golang:1.9

RUN mkdir -p /go/src/github.com/purplebooth/mnemonic
WORKDIR /go/src/github.com/purplebooth/mnemonic
RUN go get -u github.com/golang/dep/cmd/dep
RUN go get github.com/onsi/ginkgo/ginkgo
RUN go get github.com/onsi/gomega
COPY . .
RUN dep ensure
RUN (cd mnemonic && go test)
RUN go build  -o binary -ldflags "-linkmode external -extldflags -static" -a cmd/mnemonic/mnemonic.go

FROM ubuntu

WORKDIR /tmp
RUN apt-get update && apt-get install -y curl
RUN curl -o /tmp/wn3.1.dict.tar.gz http://wordnetcode.princeton.edu/wn3.1.dict.tar.gz
RUN tar xzvf /tmp/wn3.1.dict.tar.gz

FROM scratch
COPY --from=0 /go/src/github.com/purplebooth/mnemonic/binary /mnemonic
COPY --from=1 /tmp/dict /dict
ENTRYPOINT ["/mnemonic", "/dict"]
