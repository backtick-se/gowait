#!/bin/bash

# build and push images
docker compose build daemon base
docker compose build python
docker compose push daemon base python

# delete any existing daemon pod
kubectl delete pod cowaitd
sleep 1

# re-create daemon
kubectl apply -f kubernetes/daemon.yml

# grab logs
kubectl wait --for=condition=Ready pod/cowaitd
kubectl logs -f cowaitd

# clean up
kubectl delete pod cowaitd
