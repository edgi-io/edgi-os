ARG REPO
ARG TAG
FROM ${REPO}/edgi-k3s:${TAG} as k3s

ARG REPO
ARG TAG
FROM ${REPO}/edgi-bin:${TAG} as bin

ARG REPO
ARG TAG
FROM ${REPO}/edgi-base:${TAG} as base
ARG VERSION

COPY --from=k3s /output/  /output/edgi/system/k3s/
COPY --from=bin /output/  /output/edgi/system/edgi/${VERSION}/

WORKDIR /output/edgi/system/k3s
RUN mkdir -vp $(cat version) /output/sbin
RUN mv -vf crictl ctr kubectl /output/sbin/
RUN ln -sf $(cat version) current
RUN mv -vf install.sh current/k3s-install.sh
RUN mv -vf k3s current/
RUN rm -vf version *.sh
RUN ln -sf /edgi/system/k3s/current/k3s /output/sbin/k3s

WORKDIR /output/edgi/system/edgi
RUN ln -sf ${VERSION} current
RUN ln -sf /edgi/system/edgi/current/edgi /output/sbin/edgi
RUN ln -sf edgi /output/sbin/init
