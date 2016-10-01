FROM golang:1.7

RUN go get \
  github.com/kardianos/govendor \
  github.com/maxbrunsfeld/counterfeiter \
  github.com/onsi/ginkgo/ginkgo

RUN curl -o docker.tgz https://get.docker.com/builds/Linux/x86_64/docker-1.12.1.tgz && \
tar -xzf docker.tgz && \
mv ./docker/docker /usr/local/bin
