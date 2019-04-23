FROM ubuntu
MAINTAINER John Guo "unistd1999@gmail.com"
WORKDIR /data/webserver
COPY ./webserver /data/webserver/
COPY ./static /data/webserver/static
COPY ./views /data/webserver/views
COPY ./config /data/webserver/config
EXPOSE 8080
CMD ./webserver
