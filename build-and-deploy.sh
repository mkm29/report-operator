#!/bin/bash

# This script is used to build and deploy the application to the server.
set -x
set -e

img="report-operator"
tag="0.1.2"
docker_username="smigula"
docker_registry="localhost:5000"
namespace="trivy-operator"

# make generate
# make manifests
# make install

cd config/manager
kustomize edit set namespace $namespace
kustomize edit set image controller=${docker_registry}/${docker_username}/${img}:v${tag}
cp manager.yaml manager.yaml.bak
sed -e "s|_AWS_DEFAULT_REGION_|${AWS_DEFAULT_REGION}|g" -i manager.yaml
sed -e "s|_AWS_ACCESS_KEY_ID_|${AWS_ACCESS_KEY_ID}|g" -i manager.yaml
sed -e "s|_AWS_SECRET_ACCESS_KEY_|${AWS_SECRET_ACCESS_KEY}|g" -i manager.yaml
cd ../..
cd config/default
kustomize edit set namespace $namespace
cd ../..

# # If you need to change the host or username, please pass in to make
make docker-build VERSION=${tag} #docker_registry=${docker_registry} DOCKER_USERNAME=${docker_username} TAG=${tag}
make docker-tag VERSION=${tag} #IMG=${img} docker_registry=${docker_registry} DOCKER_USERNAME=${docker_username} TAG=${tag}
make docker-push VERSION=${tag} #IMG=${img} docker_registry=${docker_registry} DOCKER_USERNAME=${docker_username} TAG=${tag}

kustomize build config/default | kubectl apply -f -
#make deploy #IMG=${img} docker_registry=${docker_registry} DOCKER_USERNAME=${docker_username} TAG=${tag}

# kubectl apply -f config/samples/cache_v1alpha1_memcached.yaml
# reset manager.yaml
rm config/manager/manager.yaml
mv config/manager/manager.yaml.bak config/manager/manager.yaml