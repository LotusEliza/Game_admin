#!/bin/bash
PROJECT=strategy-dev
DBPORT=26263
ADMINPORT=26163
# put here your preferred data path
DATAPATH=/tmp/aos
DATABASENAME=strategy

mkdir -p ${DATAPATH}/data

docker stop ${PROJECT}
docker rm ${PROJECT}

docker run -d --name=${PROJECT} \
    --restart=unless-stopped \
    -p ${DBPORT}:26257 -p ${ADMINPORT}:8080 \
    -v ${DATAPATH}/data:/cockroach/cockroach-data \
    cockroachdb/cockroach:v19.1.2 start --insecure
# wait for DB up
sleep 5
docker exec -it ${PROJECT} ./cockroach sql --insecure --host="127.0.0.1" --execute="CREATE DATABASE IF NOT EXISTS ${DATABASENAME};"
