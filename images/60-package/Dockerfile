ARG REPO
ARG TAG
FROM ${REPO}/edgi-kernel:${TAG} as kernel

ARG REPO
ARG TAG
FROM ${REPO}/edgi-package:${TAG}
ARG VERSION

COPY --from=kernel /output/ /output/edgi/system/kernel/

WORKDIR /output/edgi/system/kernel
RUN mkdir -vp $(cat version)
RUN ln -sf $(cat version) current
RUN mv -vf initrd kernel.squashfs current/
RUN rm -vf version vmlinuz
