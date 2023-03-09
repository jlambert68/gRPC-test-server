# Compile stage
FROM golang:1.19 AS build-env

ADD . /dockerdev
WORKDIR /dockerdev

COPY go.* ./
RUN go mod tidy

RUN go build -o /simpleGrpcServer .


# Final stage
FROM debian:latest

RUN apt-get update \
 && apt-get install -y --force-yes --no-install-recommends \
      apt-transport-https \
      curl \
      ca-certificates \
 && apt-get clean \
 && apt-get autoremove \
 && rm -rf /var/lib/apt/lists/*

#FROM golang:1.13.8

EXPOSE 6670
#FROM golang:1.13.8
WORKDIR /

#ADD ca-certificates.crt /etc/ssl/certs/

COPY --from=build-env /simpleGrpcServer /
#Add data/ data/

#CMD ["/simpleGrpcServer"]
ENTRYPOINT ["/simpleGrpcServer"]



#// docker build -t  fenix-client-server .
#// docker run -p 5998:5998 -it  fenix-client-server
#// docker run -p 5998:5998 -it --env StartupType=LOCALHOST_DOCKER fenix-client-server

#//docker run --name fenix-client-server --rm -i -t fenix-client-server  bash