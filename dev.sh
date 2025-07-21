#!/bin/sh

CMD="ARGS='$@' make compiler" docker-compose -f docker-compose.dev.yaml run --rm go-llvm-dev
