#!/bin/bash

set -e

# build and push images
docker compose build base
docker compose build daemon cloud python
docker compose push daemon cloud base python

# delete any existing daemon pod
kubectl delete pod --ignore-not-found cowaitd cloud
sleep 1

# re-create daemon
kubectl apply -f kubernetes/cloud.yml
kubectl wait --for=condition=Ready pod/cloud

kubectl apply -f kubernetes/daemon.yml

# grab logs
kubectl wait --for=condition=Ready pod/cowaitd
kubectl logs -f cowaitd

# clean up
# kubectl delete pod cowaitd cloud
