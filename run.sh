#!/bin/bash

docker stop aos
docker rm aos

docker run -it --pull --name=aos registry.gitlab.com/finalist/strategy:dev
