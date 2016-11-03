#!/bin/sh

gcloud compute disks create --size 1GB disk-1
kubectl create -f gce-pv.yml
