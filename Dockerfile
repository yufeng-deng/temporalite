ARG REGISTRY=hub.artifactory.gcp.anz
ARG BUILDER_IMAGE=golang:1.19.2-buster
# Update debian image when new stable version comes out
ARG BASE_IMAGE_RUNTIME=gcr.io/anz-x-fabric-np-641432/base-images/base-debian11:v1.0.0

FROM ${REGISTRY}${REGISTRY:+/}${BUILDER_IMAGE} AS builder

ARG GOPROXY=https://artifactory.gcp.anz/artifactory/go

ENV GOPROXY=${GOPROXY}
ENV GO111MODULE=on
ENV GOFLAGS="-mod=vendor"

WORKDIR ${GOPATH:-/go}/src/temporalite

COPY . .

RUN go build -v -o ${GOPATH:-/go}/bin/ ${GOPATH:-/go}/src/temporalite/cmd/temporalite

FROM ${BASE_IMAGE_RUNTIME}

COPY --from=builder ${GOPATH:-/go}/bin/temporalite /
EXPOSE 7233
EXPOSE 8233

ENTRYPOINT ["/temporalite", "start", "--ephemeral", "-n", "default", "--ip" , "0.0.0.0"]
