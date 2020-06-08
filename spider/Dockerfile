FROM ubuntu
WORKDIR /data/spider
COPY ./spider_svc /data/spider/
COPY ./sign.js /data/spider/sign.js
COPY ./ca-certificates.crt /etc/ssl/certs/
EXPOSE 19999
ENTRYPOINT [ "././spider_svc" ]
