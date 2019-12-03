#!/bin/bash
PROJECT=strategy-dev
docker exec -it $PROJECT ./cockroach sql --insecure
