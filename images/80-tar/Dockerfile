ARG REPO
ARG TAG
FROM ${REPO}/edgi-package:${TAG} as package

ARG REPO
ARG TAG
FROM ${REPO}/edgi-base:${TAG} as base
ARG VERSION

COPY --from=package /output/   /usr/src/${VERSION}/
WORKDIR /output
RUN tar cvf userspace.tar -C /usr/src ${VERSION}
