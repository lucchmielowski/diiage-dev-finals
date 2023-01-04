#!/bin/bash
docker build -t luskidotme/catalog-service:v1 ../catalog-service/
docker build -t luskidotme/presentation-service:v1 ../presentation-service/

docker push luskidotme/catalog-service:v1
docker push luskidotme/presentation-service:v1
