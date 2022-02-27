ARG REPO
ARG TAG
FROM ${REPO}/edgi-package:${TAG} as package

ARG REPO
ARG TAG
FROM ${REPO}/edgi-base:${TAG} as base
ARG VERSION
ARG ARCH
RUN apk add xorriso grub grub-efi mtools libvirt qemu-img
RUN if [ "$ARCH" == "amd64" ]; then \
        apk add qemu-system-x86_64 grub-bios ovmf \
    ;elif [ "$ARCH" == "arm64" ]; then \
        apk add qemu-system-aarch64 \
    ;fi
RUN ln -s /usr/bin/qemu-system-* /usr/bin/qemu-system
RUN qemu-img create -f qcow2 /hd.img 40G
COPY run-kvm.sh /usr/bin/
COPY grub.cfg /usr/src/iso/boot/grub/grub.cfg

COPY --from=package /output/ /usr/src/iso/

COPY config.yaml /usr/src/iso/edgi/system/
RUN mkdir -p /output && \
    grub-mkrescue -o /output/edgi.iso /usr/src/iso/. -- -volid EDGI -joliet on && \
    [ -e /output/edgi.iso ] # grub-mkrescue doesn't exit non-zero on failure

CMD ["run-kvm.sh"]
