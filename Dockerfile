FROM alpine:latest AS userland 
ARG PUID=10001
ARG PGID=10002

RUN addgroup -g ${PGID} abc && \
    adduser -u ${PUID} -G abc nonroot --no-create-home --disabled-password
#https://wiki.alpinelinux.org/wiki/Setting_up_a_new_user

#https://medium.com/@lizrice/non-privileged-containers-based-on-the-scratch-image-a80105d6d341
#FROM ubuntu:latest AS userland 
#RUN useradd --user-group --uid 10001 scratchuser

FROM scratch
#ENV PATH=/bin:/usr/local/bin/
COPY --from=userland /etc/passwd /etc/passwd
COPY --from=userland /etc/group /etc/group
COPY --chown=nonroot:abc cuc /usr/local/bin/cuc
USER nonroot
#WORKDIR /usr/local/bin/ # dist/${BuildID}_${BuildTarget}
ENTRYPOINT [ "/usr/local/bin/cuc" ]