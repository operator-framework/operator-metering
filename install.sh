#!/bin/bash

# Modify a Tectonic install to use the new kube-state-metrics that has node ID working
kubectl -n tectonic-system patch deployment kube-state-metrics -p '{"spec":{"template":{"spec":{"containers":[{"name":"kube-state-metrics","image":"quay.io/dan_gillespie/kube-state-metrics:v0.5.0"}]}}}}'

# install query layer (optional for collection)
kubectl apply -f ./hdfs
sleep 5

kubectl apply -f ./hive
kubectl apply -f ./presto
kubectl apply -f ./promsum
kubectl apply -f prom.yaml
kubectl apply -f kube-state-metrics.yaml
