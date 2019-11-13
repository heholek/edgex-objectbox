#!/usr/bin/env bash
set -euo pipefail

# Checks out an upstream repo and build adocker image for it

if [ "$#" -ne 5 ]; then
    echo "usage: $(basename "$0") dirname commit url DOCKER_TAG_GIT_SHA DOCKER_TAG_VERSION"
fi

name=$1
commit=$2
url=$3

./build/get-upstream-repo.sh "$name" "$commit" "$url"

git_sha=$4
tag=$5

docker build \
		--label "git_sha=$git_sha" \
		-t "objectboxio/edgex-$name:$git_sha" \
		-t "objectboxio/edgex-$name:$tag" \
		build/checkouts/$name
