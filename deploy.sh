#!/bin/bash
docker compose build daemon base
docker compose build python
docker compose push daemon base python

kubectl delete pod cowaitd gowait-task
kubectl apply -f kubernetes/daemon.yml

sleep 5
kubectl logs -f cowaitd