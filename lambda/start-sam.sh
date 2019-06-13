#!/bin/bash

sam local start-api \
  --docker-volume-basedir "${VOLUME}/" \
  --docker-network aws-local \
  --host 0.0.0.0 \
  --template template.yaml
