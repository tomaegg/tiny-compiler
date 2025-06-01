#!/bin/sh

case "$1" in
"build")
  docker buildx build -t tiny-compiler:latest -f docker/Dockerfile .
  ;;
"parser" | "symtable" | "ir" | "dot")
  mkdir -p output
  docker run \
    -v "$(pwd)/example:/runtime/example" \
    -v "$(pwd)/output:/runtime/output" \
    tiny-compiler:latest "$@"
  ;;
*)
  echo "Usage: $0 [build|parser|symtable|ir|dot] [args...]"
  exit 1
  ;;
esac
