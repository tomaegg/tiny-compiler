# Tiny Compiler

## Dev Toolchain

golang: 1.24.2 with cgo

## Run

### 获取镜像

```shell
# 导入镜像
./run.sh import

# 或者构建镜像
./run.sh build

```

### 运行

- Parser

```shell
./run.sh parser scopes.rs # scopes.rs位于example目录中(挂载)
```

- Symbol Table

```shell
./run.sh symtable scopes.rs # scopes.rs位于example目录中(挂载)
```

- Visualize Symbol Table with Dot Lanugage

```shell
./run.sh dot scopes.rs # scopes.rs位于example目录中(挂载), 在output中查看产物
```

- Generate LLVM IR

```shell
./run.sh ir func.rs # func.rs位于example目录中(挂载)
```
