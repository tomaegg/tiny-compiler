# Tiny Compiler

## Dev Toolchain

golang: 1.24.2

## Run

```shell

# 构建镜像
./run.sh build

# 运行
./run.sh dot scopes.rs # scopes.rs位于example目录中(挂载), 在output中查看产物
./run.sh parser scopes.rs # scopes.rs位于example目录中(挂载)
./run.sh symtable scopes.rs # scopes.rs位于example目录中(挂载)
./run.sh ir func.rs # func.rs位于example目录中(挂载)

```
