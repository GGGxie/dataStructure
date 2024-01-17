FROM golang:1.21 AS build

ENV GOPROXY="https://goproxy.cn,direct"
ENV GO111MODULE="on"
WORKDIR /app/dataStructure
COPY . /app/dataStructure/

RUN go mod download
