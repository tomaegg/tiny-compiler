FROM golang:1.24-bookworm AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends \
  llvm-19 llvm-19-dev make && \
  apt-get clean && rm -rf /var/lib/apt/lists/*

# use go build cache
COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn,direct && go mod download
COPY . .

# paralle build with make
RUN make --file docker/Makefile --jobs $(nproc)

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
