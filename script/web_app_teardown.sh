#!/bin/sh
kubectl delete deployment webapp
kubectl delete services webapp-service
sleep 3