#!/usr/bin/env bash
git push
tag=v.$(date +'%Y%m%d.%H%M').$(git rev-parse --short HEAD)
git tag "$tag"
git push origin "$tag"
exit
