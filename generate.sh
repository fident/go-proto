#!/bin/bash
FILES=vendor/github.com/fident/proto/*.proto
for f in $FILES
do
  file=$(basename "$f")
  package="${file%.*}"
  echo "Processing $package $f"
  mkdir -p $package
  protoc -I vendor/github.com/fident/proto --go_out=plugins=grpc:$package $f
  git add $package/*
done
