#!/bin/bash
docker build -t registry.gitlab.com/finalist/strategy:dev .
docker push registry.gitlab.com/finalist/strategy:dev
