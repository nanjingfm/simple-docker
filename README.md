# DEMO

通过几十行 go 代码实现一个简易的 docker。

> **注意** 代码只能在 linux 环境运行。

先执行 sh gen-rootfs.sh 生成一个最简单的运行时目录。

然后编译 go 代码，生成一个可执行文件。

```shell
go build -o gorun .
```

执行以下命令就可以切换到容器运行时环境了。

```shell
./gorun run /bin/sh
```
