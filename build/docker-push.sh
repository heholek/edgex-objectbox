#!/usr/bin/env bash
set -euo pipefail

# Note: to delete a manifest, remove it manually from ~/.docker/manifests/

version=$(cat ./VERSION)

images=$(docker images "objectboxio/edgex-*:$(uname -m)-$version" --format "{{.Repository}}:{{.Tag}}")

echo "$images"

read -p "Would you like to push all images? " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo "$images" | xargs -n1 docker push
fi
