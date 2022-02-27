ARG REPO
ARG TAG
FROM ${REPO}/edgi-bin:${TAG} as bin

FROM ${REPO}/edgi-kernel-stage1:${TAG} as kernel

FROM ${REPO}/edgi-base:${TAG}
ARG TAG
RUN apk add squashfs-tools
COPY --from=kernel /output/ /usr/src/kernel/

RUN mkdir -p /usr/src/initrd/lib && \
    cd /usr/src/kernel && \
    tar cf - -T initrd-modules -T initrd-firmware | tar xf - -C /usr/src/initrd/ && \
    depmod -b /usr/src/initrd $(cat /usr/src/kernel/version)

RUN mkdir -p /output && \
    cd /usr/src/kernel && \
    depmod -b . $(cat /usr/src/kernel/version) && \
    mksquashfs . /output/kernel.squashfs

RUN cp /usr/src/kernel/version /output/ && \
    cp /usr/src/kernel/vmlinuz /output/

COPY --from=bin /output/ /usr/src/edgi/
RUN cd /usr/src/initrd && \
    mkdir -p edgi/system/edgi/${TAG} && \
    cp /usr/src/edgi/edgi edgi/system/edgi/${TAG} && \
    ln -s ${TAG} edgi/system/edgi/current && \
    ln -s /edgi/system/edgi/current/edgi init
    
RUN cd /usr/src/initrd && \
    find . | cpio -H newc -o | gzip -c -1 > /output/initrd
