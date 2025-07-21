#!/bin/sh

echo "running [$@]"

CMD="$@" docker-compose -f docker-compose.dev.yaml run --rm go-llvm-dev

