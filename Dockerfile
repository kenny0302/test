FROM golang:latest
WORKDIR /test
ADD . /test
RUN cd /test && go build
EXPOSE 8000
ENTRYPOINT ./main