FROM alpine:latest as userland 
RUN adduser -u 10001 scratchuser

#https://medium.com/@lizrice/non-privileged-containers-based-on-the-scratch-image-a80105d6d341
#FROM ubuntu:latest
#RUN useradd -u 10001 scratchuser

FROM scratch
#ENV PATH=/bin
COPY cuc /usr/local/bin/cuc
COPY --from=userland /etc/passwd /etc/passwd
USER scratchuser
ENTRYPOINT [ "/usr/local/bin/cuc" ]