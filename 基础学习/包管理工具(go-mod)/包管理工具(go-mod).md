## 一个包管理工具应该有以下功能：

### 基本功能
- 依赖管理
- 依赖包版本控制
- 对应的包管理平台
- 可以私有化部署

### 加分
- 代码包是否可以复用
- 构建，测试,打包
- 发布上线

对比上面几点：来谈谈今天主角是go mod。

### 使用go path问题
1. 代码开发必须在go path src目录下，不然，就有问题
2. 依赖手动管理
3. 依赖包没有版本可言

### go mod介绍
go modules 是 golang 1.11 新加的特性。现在1.12 已经发布了，是时候用起来了。Modules官方定义为：

```text
模块是相关Go包的集合。modules是源代码交换和版本控制的单元。go命令直接支持使用modules，包括记录和
解析对其他模块的依赖性。modules替换旧的基于GOPATH的方法来指定在给定构建中使用哪些源文件。
```

### 如何使用go mod
首先，必须升级go到1.11,目前版本是1.14，下面我以我自己升级演示：

```shell
### 卸载旧版本，删除对应文件
brew uninstall -f go

### 更新一下brew

brew update


### 安装go
brew install go
```

上面升级完了，使用 go version看下版本
```shell
go version go1.14.1 darwin/amd64
```

下面设置go mod和go proxy
```bash
go env -w GOBIN=/Users/youdi/go/bin
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct // 使用七牛云的
```
注意：go env -w会将配置写到 GOENV="/Users/youdi/Library/Application Support/go/env"

### GO111MODULE
GO111MODULE 有三个值：off, on和auto（默认值）。

GO111MODULE=off，go命令行将不会支持module功能，寻找依赖包的方式将会沿用旧版本那种通过vendor目录或者GOPATH模式来查找。

GO111MODULE=on，go命令行会使用modules，而一点也不会去GOPATH目录下查找。

GO111MODULE=auto，默认值，go命令行将会根据当前目录来决定是否启用module功能。这种情况下可以分为两种情形：

```text
当前目录在GOPATH/src之外且该目录包含go.mod文件
当前文件在包含go.mod文件的目录下面。
```

当modules功能启用时，依赖包的存放位置变更为$GOPATH/pkg，允许同一个package多个版本并存，且多个项目可以共享缓存的 module

我们看下目录：
cd /Users/youdi/go/pkg
```text
├── darwin_amd64
│   ├── github.com
│   ├── go.etcd.io
│   ├── golang
│   ├── golang.org
│   ├── gopkg.in
│   ├── quickstart
│   └── uc.a
├── mod
│   ├── cache
│   ├── github.com
│   ├── golang.org
│   ├── google.golang.org
│   └── gopkg.in
└── sumdb
    └── sum.golang.org
```

### go mod命令
golang 提供了 go mod命令来管理包。
go help mod

```text
Go mod provides access to operations on modules.

Note that support for modules is built into all the go commands,
not just 'go mod'. For example, day-to-day adding, removing, upgrading,
and downgrading of dependencies should be done using 'go get'.
See 'go help modules' for an overview of module functionality.

Usage:

    go mod <command> [arguments]

The commands are:

    download    download modules to local cache
    edit        edit go.mod from tools or scripts
    graph       print module requirement graph
    init        initialize new module in current directory
    tidy        add missing and remove unused modules
    vendor      make vendored copy of dependencies
    verify      verify dependencies have expected content
    why         explain why packages or modules are needed

Use "go help mod <command>" for more information about a command.
```
go mod 有以下命令：
![img.png](https://pic1.zhimg.com/80/v2-d83c2d54165a9090376453e68220a08c_720w.webp)

比较常用的是 init,tidy, edit

### 使用go mod管理一个新项目

