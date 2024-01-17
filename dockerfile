FROM golang:1.21 AS build

ENV GOPROXY="https://goproxy.cn,direct"
ENV GO111MODULE="on"
WORKDIR /app/dataStructure
COPY . /app/dataStructure/

RUN go mod download
# RUN CGO_ENABLED=0 GOOS=linux go build -o bk-api-adapter ./main.go

# FROM alpine:latest
# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
# RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
#     && echo "Asia/Shanghai" > /etc/timezone \
#     && apk del tzdata
# WORKDIR /
# COPY --from=build /app/dataStructure .

EXPOSE 30000
# ENTRYPOINT ["/bk-api-adapter"]