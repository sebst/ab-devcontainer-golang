#!/bin/env bash

readonly package=${1:-default}
readonly version=${2:-latest}

echo "Installing $package version $version" 

for i in {1..5} ; do
    sleep 1
    echo -e ".\n\c"
done