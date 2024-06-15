#!/bin/sh

curl -O https://www.busybox.net/downloads/binaries/1.35.0-x86_64-linux-musl/busybox
chmod +x busybox
mkdir -p ./rootfs/{bin,sbin,etc,proc,sys,usr/{bin,sbin}}

# 复制 busybox 二进制文件
cp busybox ./rootfs/bin

# 创建必要的符号链接
cd ./rootfs/bin
for cmd in $(./busybox --list); do
  ln -s busybox $cmd
done
