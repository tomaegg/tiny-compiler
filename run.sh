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
  DOCKER_BUILDKIT=1 docker build -t "$IMAGE_NAME" -f Dockerfile.runtime .
  ;;
"compile")
  shift 1
  mkdir -p output
  docker run --rm -t \
    -v "$(pwd):/runtime" \
    tiny-compiler:latest "$@"
  ;;
*)
  echo "Usage: $0 [export|import|build|compile] [args...]"
  exit 1
  ;;
esac
