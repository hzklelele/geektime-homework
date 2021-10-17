FROM ubuntu
ENV MY_SERVICE_PORT=80
LABEL multi.label1="value1"
ADD bin/amd64/httpserver /httpserver
EXPOSE 80
ENTRYPOINT /httpserver