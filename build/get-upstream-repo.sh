#!/usr/bin/env bash
set -euo pipefail

# Checks out an upstream repo - we build them manually (necessary for ARM)

if [ "$#" -ne 3 ]; then
    echo "usage: $(basename "$0") dirname commit url"
fi

name=$1
commit=$2
url=$3

dir="build/checkouts/$name"

# make a new blank repository
rm -rf "$dir"
git init "$dir"
cd "$dir"

# add a remote
git remote add origin $url

# fetch a commit (or branch or tag) of interest
# Note: the full history up to this commit will be retrieved unless
#       you limit it with '--depth=...' or '--shallow-since=...'
git fetch origin "$commit" --depth=1

# reset this repository's master branch to the commit of interest
git reset --hard FETCH_HEAD

