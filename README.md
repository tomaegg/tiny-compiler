# Tiny Compiler

## Dev Toolchain

golang: 1.24.2 with cgo

## Run

### 获取镜像

```shell
# 如果有本地镜像, 导入镜像
./run.sh import

# 或者构建镜像
./run.sh build

```

### 运行

- Visualize Symbol Table with Dot Lanugage

```shell
./run.sh compile example/scopes.rs --stage=semantic --visualize # func.rs位于example目录中(挂载)
```

- Generate LLVM IR

```shell
./run.sh compile example/func.rs --stage=ir # func.rs位于example目录中(挂载)
```
