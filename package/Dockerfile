FROM scratch
COPY build/edgi/system/ /edgi/system/
ENV PATH /edgi/system/edgi/current:/edgi/system/k3s/current:${PATH}
ENTRYPOINT ["edgi"]
CMD ["help"]
