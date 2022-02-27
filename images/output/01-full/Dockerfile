ARG REPO
ARG TAG
FROM ${REPO}/edgi-tar:${TAG} as tar

ARG REPO
ARG TAG
FROM ${REPO}/edgi-iso:${TAG} as iso

ARG REPO
ARG TAG
FROM ${REPO}/edgi-kernel:${TAG} as kernel

ARG REPO
ARG TAG
FROM ${REPO}/edgi-base:${TAG}
ARG ARCH

COPY --from=kernel /output/vmlinuz /output/edgi-vmlinuz-${ARCH}
COPY --from=kernel /output/initrd /output/edgi-initrd-${ARCH}
COPY --from=kernel /output/kernel.squashfs /output/edgi-kernel-${ARCH}.squashfs
COPY --from=kernel /output/version /output/edgi-kernel-version-${ARCH}
COPY --from=iso /output/edgi.iso /output/edgi-${ARCH}.iso
COPY --from=tar /output/userspace.tar /output/edgi-rootfs-${ARCH}.tar
RUN gzip /output/edgi-rootfs-${ARCH}.tar
