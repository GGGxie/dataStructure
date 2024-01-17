FROM golang:1.21


WORKDIR /app/dataStructure
COPY . /app/dataStructure/

RUN echo "export PATH=/usr/local/go/bin:$PATH" && echo "export GO111MODULE=on" && echo "export GOPROXY=https://goproxy.cn,direct"
RUN go mod download

