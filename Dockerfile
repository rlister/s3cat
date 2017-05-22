FROM scratch
MAINTAINER Ric Lister <rlister@gmail.com>

ADD certs/ca-certificates.crt /etc/ssl/certs/
ADD s3cat /

ENTRYPOINT [ "/s3cat" ]