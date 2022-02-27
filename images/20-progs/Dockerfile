ARG REPO
ARG TAG
ARG VERSION
FROM ${REPO}/edgi-gobuild:${TAG} as gobuild

ENV LINUXKIT v0.8

FROM gobuild as linuxkit
ENV GO111MODULE off
RUN git clone https://github.com/linuxkit/linuxkit.git $GOPATH/src/github.com/linuxkit/linuxkit
WORKDIR $GOPATH/src/github.com/linuxkit/linuxkit/pkg/metadata
RUN git checkout -b current $LINUXKIT
RUN gobuild -o /output/metadata

FROM gobuild as edgi
ARG VERSION
COPY go.mod $GOPATH/src/github.com/edgi-io/edgi-os/
COPY go.sum $GOPATH/src/github.com/edgi-io/edgi-os/
COPY /pkg/ $GOPATH/src/github.com/edgi-io/edgi-os/pkg/
COPY /main.go $GOPATH/src/github.com/edgi-io/edgi-os/
COPY /vendor/ $GOPATH/src/github.com/edgi-io/edgi-os/vendor/
WORKDIR $GOPATH/src/github.com/edgi-io/edgi-os
RUN gobuild -mod=readonly -o /output/edgi

FROM gobuild
COPY --from=linuxkit /output/ /output/
COPY --from=edgi /output/ /output/
WORKDIR /output
RUN git clone --branch v0.7.0 https://github.com/ahmetb/kubectx.git \
 && chmod -v +x kubectx/kubectx kubectx/kubens
