#!/bin/bash

AppName="smartvoice-platform"
IMAGE_NAME="applet-service"

user="annotation"
password="7GXZUsYFYP.Jrq"
CI_COMMIT_TAG=`git log --pretty=format:"%h" -1`
DOCKER_REGISTRY_HOST="harbor.cloudminds.com"
VERSION="v1.0.0"


docker build  --no-cache -t harbor.cloudminds.com/$AppName/$IMAGE_NAME:$VERSION.$CI_COMMIT_TAG .
echo DOCKER_REGISTRY_USER=$user DOCKER_REGISTRY_PASSWORD=$password
echo $password |  docker login -u $user --password-stdin $DOCKER_REGISTRY_HOST >/dev/null 2>&1 && docker push harbor.cloudminds.com/$AppName/$IMAGE_NAME:$VERSION.$CI_COMMIT_TAG
