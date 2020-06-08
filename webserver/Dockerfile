FROM ubuntu
WORKDIR /data/webserver
COPY ./web /data/webserver/web
COPY ./static /data/webserver/static
COPY ./views /data/webserver/views
# COPY ./config /data/webserver/config
VOLUME [ "/data/webserver/config" ]
EXPOSE 8080
ENTRYPOINT [ "./web" ] 
