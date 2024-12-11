#! /usr/bin/env bash
set -e

TAG=$(date +%y%m%d)
TARGETS="sgx-runtime sgx-sdk sgx-go"

for target in $TARGETS; do
    echo "Building: $target"
    docker build --target $target -t nithinkatla00/$target:latest -t nithinkatla00/$target:$TAG -f Dockerfile.base .
    docker push nithinkatla00/$target
    echo "Build complete: $target"
done

