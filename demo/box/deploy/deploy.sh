#!/bin/bash

# 定义一些变量
DOCKERFILE_PATH="./deploy/deploy.dockerfile"  # Dockerfile 的路径
DOCKER_IMAGE_NAME="box"       # 镜像名称
DOCKER_IMAGE_TAG="v1.0"         # 镜像标签
REGISTRY="15992686641"              # 镜像仓库地址（如果有）

# 构建 Docker 镜像
echo docker build -t $DOCKER_IMAGE_NAME:$DOCKER_IMAGE_TAG -f $DOCKERFILE_PATH .
docker build -t $DOCKER_IMAGE_NAME:$DOCKER_IMAGE_TAG -f $DOCKERFILE_PATH .

# 如果您有一个自定义的镜像仓库，请登录（如果需要）
# docker login $REGISTRY

# 推送 Docker 镜像到仓库
echo docker push $DOCKER_IMAGE_NAME:$DOCKER_IMAGE_TAG
docker push $DOCKER_IMAGE_NAME:$DOCKER_IMAGE_TAG

# 如果使用了自定义仓库，还需要标记并推送到自定义仓库
if [ -n "$REGISTRY" ]; then
  echo   docker tag $DOCKER_IMAGE_NAME:$DOCKER_IMAGE_TAG $REGISTRY/$DOCKER_IMAGE_NAME:$DOCKER_IMAGE_TAG
  docker tag $DOCKER_IMAGE_NAME:$DOCKER_IMAGE_TAG $REGISTRY/$DOCKER_IMAGE_NAME:$DOCKER_IMAGE_TAG
  echo   docker push $REGISTRY/$DOCKER_IMAGE_NAME:$DOCKER_IMAGE_TAG
  docker push $REGISTRY/$DOCKER_IMAGE_NAME:$DOCKER_IMAGE_TAG
fi

# 清理本地镜像（可选）
echo docker rmi $REGISTRY/$DOCKER_IMAGE_NAME:$DOCKER_IMAGE_TAG
docker rmi $REGISTRY/$DOCKER_IMAGE_NAME:$DOCKER_IMAGE_TAG