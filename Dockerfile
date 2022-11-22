ARG BUILDER_IMAGE=golang
ARG BUILDER_VERSION=1.18-alpine

FROM ${BUILDER_IMAGE}:${BUILDER_VERSION} AS builder

WORKDIR /go/src/chaincode
COPY go.mod go.sum ./

COPY . .
RUN CGO_ENABLED=0 go build -v -o /go/bin/chaincode

FROM alpine:3.15
COPY --chown=65534 --from=builder /go/bin/chaincode /
USER 65534

ENTRYPOINT ["/chaincode"]

