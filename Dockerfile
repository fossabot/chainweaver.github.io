FROM golang:1.11.5-alpine3.8 AS Build

ENV GOLANG_PROTOBUF_VERSION v1.2.0
ENV GRPC_GATEWAY_VERSION v1.5.1
ENV PROTOC_GEN_DOC_VERSION v1.1.0

RUN apk add --update \
  && apk --no-cache add git

# Install dep
RUN go get -u github.com/golang/dep/cmd/dep

# Install protobuf for golang
RUN mkdir -p $GOPATH/src/github.com/golang
WORKDIR $GOPATH/src/github.com/golang
RUN git clone https://github.com/golang/protobuf.git
WORKDIR $GOPATH/src/github.com/golang/protobuf
RUN git checkout -b $GOLANG_PROTOBUF_VERSION refs/tags/$GOLANG_PROTOBUF_VERSION
WORKDIR $GOPATH/src/github.com/golang/protobuf/protoc-gen-go
RUN go install

# Install grpc-gateway and swagger
WORKDIR $GOPATH/src/github.com/grpc-ecosystem
RUN git clone https://github.com/grpc-ecosystem/grpc-gateway.git
WORKDIR $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway
RUN git checkout -b $GRPC_GATEWAY_VERSION refs/tags/$GRPC_GATEWAY_VERSION
RUN dep ensure
WORKDIR $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
RUN go install
WORKDIR $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
RUN go install

# Install protobuf documentation generator
WORKDIR $GOPATH/src/github.com/pseudomuto
RUN git clone https://github.com/pseudomuto/protoc-gen-doc.git
WORKDIR $GOPATH/src/github.com/pseudomuto/protoc-gen-doc
RUN git checkout -b $PROTOC_GEN_DOC_VERSION refs/tags/$PROTOC_GEN_DOC_VERSION
RUN dep ensure
WORKDIR $GOPATH/src/github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
RUN go install


FROM alpine:3.8

RUN apk add --update \
  && apk --no-cache add protobuf protobuf-dev jq py-pip bash

# Install yq
RUN pip install --upgrade pip yq

COPY --from=build /go/bin/protoc-gen-go /usr/bin/
COPY --from=build /go/bin/protoc-gen-grpc-gateway /usr/bin/
COPY --from=build /go/bin/protoc-gen-swagger /usr/bin/
COPY --from=build /go/bin/protoc-gen-doc /usr/bin/
COPY --from=build /go/src/github.com/grpc-ecosystem/grpc-gateway /src/github.com/grpc-ecosystem/grpc-gateway
COPY --from=build /go/src/github.com/golang/protobuf /src/github.com/golang/protobuf

WORKDIR /go/src/github.com/fairway-corp/blockchain-protobuf
ENTRYPOINT [ "/bin/bash", "genproto.sh" ]