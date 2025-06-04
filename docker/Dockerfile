FROM golang:1.24-bookworm AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends \
  llvm-19 llvm-19-dev && \
  apt-get clean && rm -rf /var/lib/apt/lists/*

COPY . /app

RUN mkdir -p build && \
  go build -o build/parser cmd/parser/main.go && \
  go build -o build/symtable cmd/symtable/main.go && \
  go build -tags=llvm19 -o build/ir cmd/ir/main.go

RUN strip build/*

FROM debian:bookworm-slim AS runtime

WORKDIR /runtime

# timezone
ENV TZ=Asia/Shanghai

# llvm19, graphviz
RUN apt-get update && apt-get install -y --no-install-recommends \
  llvm-19-runtime graphviz && apt-get clean && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/build/* /runtime/

COPY docker/entrypoint.sh /runtime/

RUN chmod +x /runtime/entrypoint.sh

ENTRYPOINT ["/runtime/entrypoint.sh"]
