#!/bin/sh

IMAGE_NAME=tiny-compiler:latest
EXPORT_IMAGE=tiny-compiler-latest.tar

case "$1" in
"import")
  docker load -i "$EXPORT_IMAGE"
  ;;
"export")
  docker save -o "$EXPORT_IMAGE" "$IMAGE_NAME"
  ;;
"build")
  DOCKER_BUILDKIT=1 docker build -t "$IMAGE_NAME" -f docker/Dockerfile .
  ;;
"parser" | "symtable" | "ir" | "dot")
  mkdir -p output
  docker run \
    -v "$(pwd)/example:/runtime/example" \
    -v "$(pwd)/output:/runtime/output" \
    tiny-compiler:latest "$@"
  ;;
*)
  echo "Usage: $0 [export|import|build|parser|symtable|ir|dot] [args...]"
  exit 1
  ;;
esac
