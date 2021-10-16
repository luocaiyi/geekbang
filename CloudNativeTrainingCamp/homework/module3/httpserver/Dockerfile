FROM ubuntu
ENV MY_SERVICE_PORT=80
LABEL multi.lable1="value1" other="value2"
ADD bin/amd64/httpserver /httpserver
EXPOSE 80
ENTRYPOINT /httpserver
