FROM golang:1.6.2-alpine

RUN apk -U add curl git && \
curl -o docker.tgz https://get.docker.com/builds/Linux/x86_64/docker-1.11.1.tgz && \
tar -xzf docker.tgz && \
mv ./docker/docker /usr/local/bin
