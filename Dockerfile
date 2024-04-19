# Build binary using golang:latest
FROM golang:1.21-alpine3.18 as builder
RUN mkdir -p /tmp/log-generator
COPY ./ /tmp/log-generator
RUN ls -alh /tmp/log-generator
WORKDIR /tmp/log-generator/
RUN GOOS=linux GOARCH=amd64 go build -v -o .

# Copy binary and configurations
FROM centos:8
USER root
RUN rm -rf /etc/yum.repos.d/*
# COPY ./repos/*.repo /etc/yum.repos.d/
RUN bash -c "$(curl -L https://setup.vector.dev)"
RUN yum install -y vector
RUN mkdir -p /var/log/nginx/
RUN mkdir -p /scripts/
COPY ./vector/vector.yaml /etc/vector/
COPY ./scripts/entrypoint.sh /scripts/
RUN chmod +x /scripts/entrypoint.sh
COPY --from=builder /tmp/log-generator/log-generator /usr/bin/
EXPOSE 9090

ENTRYPOINT [ "/scripts/entrypoint.sh" ]