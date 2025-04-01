# Hamster-Go-Framework

### 框架介绍

土拨鼠Golang后端服务框架，主要为后台服务程序设计，集成了单应用程序多服务启动指令、动态配置、日志、单元测试、Redis缓存、Mysql数据库、通用库等组件等等。 

### 目录结构

```
cmd     配置程序中各个子服务启动指令
common  业务逻辑相关的代码库
core    业务逻辑无关的通用代码库
log     系统日志记录组件
servers 各个子服务应用程序
test    单元测试
main.yaml 系统配置，修给后动态加载内存
```

### 使用方式

```

启动API服务：go run . api
启动结算服务：go run . settle
加密系统密钥：go run . encrypt "要加密的字符串"
启动API测试用例：go run ./test/api-token-list

修改文件版本号：VERSION
编译LINUX执行文件：./build.sh

```

