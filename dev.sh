#!/bin/sh


MakeCMD=$(printf "ARGS='%s' make compiler" "$*")

CMD="$MakeCMD" docker-compose -f docker-compose.dev.yaml run --rm go-llvm-dev
