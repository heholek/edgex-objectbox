#!/usr/bin/env bash
set -euo pipefail

# Note: to delete a manifest, remove it manually from ~/.docker/manifests/

version=$(cat ./VERSION)

declare -a archs=("x86_64" "armv7l")

# exact tag for the current architecture & version - used to find the list of images
thisArchTag=$(uname -m)-$version

# get the list of image names (without tags) to combine
images=$(docker image ls | grep "objectboxio/edgex-" | grep "$thisArchTag" | cut -d' ' -f1)

declare -a creates=()
declare -a pushes=()
for image in ${images[*]}; do
    target="$image:$version"

    manifestCreate="docker manifest create $target "
    for arch in ${archs[*]}; do
        manifestCreate="$manifestCreate $image:$arch-$version"
    done
    creates+=("$manifestCreate")
    pushes+=("docker manifest push $target")
done

printf '%s\n' "${creates[@]}"
printf '%s\n' "${pushes[@]}"