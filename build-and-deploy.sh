#!/bin/bash

# This script is used to build and deploy the application to the server.
set -x
set -e

IMAGE_NAME="report-operator"
IMAGE_TAG="0.1.15"
DOCKER_USERNAME="smigula"
DOCKER_REGISTRY="localhost:5000"
NAMESPACE="trivy-operator"
AWS_REGION=${AWS_DEFAULT_REGION}
AWS_PROFILE=trivy-operator-profile
S3_BUCKET="mitchmurphy-trivy-report"
# extract access key from ~./aws/credentials
# AWS_ACCESS_KEY_ID=$(grep aws_access_key_id ~/.aws/credentials | awk '{print $3}')

# make generate
# make manifests
# make install

cd config/manager
kustomize edit set namespace $NAMESPACE
kustomize edit set image controller=${DOCKER_REGISTRY}/${DOCKER_USERNAME}/${IMAGE_NAME}:v${IMAGE_TAG}
cp manager.yaml manager.yaml.bak
# either set these variables in this file or it will default to using those from the environment
sed -e "s|_AWS_DEFAULT_REGION_|${AWS_REGION}|g" -i manager.yaml
sed -e "s|_AWS_ACCESS_KEY_ID_|${AWS_ACCESS_KEY_ID}|g" -i manager.yaml
sed -e "s|_AWS_SECRET_ACCESS_KEY_|${AWS_SECRET_ACCESS_KEY}|g" -i manager.yaml
sed -e "s|_S3_BUCKET_|${S3_BUCKET}|g" -i manager.yaml

cd ../..
cd config/default
kustomize edit set namespace $NAMESPACE
cd ../..

# For ECR
ECR_REGISTRY="${AWS_ACCOUNT}.dkr.ecr.${AWS_REGION}.amazonaws.com"
# make ecr-login AWS_PROFILE=$AWS_PROFILE AWS_REGION=$AWS_REGION ECR_REGISTRY=$ECR_REGISTRY

# If you need to change the host or username, please pass in to make
make docker-build VERSION=${IMAGE_TAG}
make docker-tag VERSION=${IMAGE_TAG}
make docker-push VERSION=${IMAGE_TAG}

kustomize build config/default | kubectl apply -f -
#make deploy #IMG=${img} docker_registry=${docker_registry} DOCKER_USERNAME=${docker_username} TAG=${tag}

# kubectl apply -f config/samples/cache_v1alpha1_memcached.yaml
# reset manager.yaml
rm config/manager/manager.yaml
mv config/manager/manager.yaml.bak config/manager/manager.yaml