ARG REPO
ARG TAG
FROM ${REPO}/edgi-tar:${TAG} as tar

ARG REPO
ARG TAG
FROM ${REPO}/edgi-base:${TAG}
ARG ARCH

COPY --from=tar /output/userspace.tar /output/edgi-rootfs-${ARCH}.tar
RUN gzip /output/edgi-rootfs-${ARCH}.tar
