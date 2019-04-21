FROM ubuntu
MAINTAINER John Guo "unistd1999@gmail.com"
WORKDIR /data/spider
COPY ./spider_svc /data/spider/
COPY ./sign.js /data/spider/sign.js
COPY ./ca-certificates.crt /etc/ssl/certs/
EXPOSE 19999
CMD ./spider_svc
