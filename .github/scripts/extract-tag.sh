#!/bin/bash

TAG=${GITHUB_REF/refs\/tags|heads\//}

if [ "$TAG" = "main" ]
then
  echo ::set-output name=DOCKER_IMAGE_TAG::latest
else
  echo ::set-output name=DOCKER_IMAGE_TAG::"$TAG"
fi