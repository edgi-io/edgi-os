ARG REPO
ARG TAG
FROM ${REPO}/edgi-rootfs:${TAG} as rootfs

ARG REPO
ARG TAG
FROM ${REPO}/edgi-progs:${TAG} as progs

ARG REPO
ARG TAG
FROM ${REPO}/edgi-base:${TAG}

COPY --from=rootfs /output/rootfs.squashfs /usr/src/
COPY install.sh /output/edgi-install.sh
COPY --from=progs /output/edgi /output/edgi
RUN echo -n "_sqmagic_" >> /output/edgi
RUN cat /usr/src/rootfs.squashfs >> /output/edgi
