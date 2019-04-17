# gomicroservices
微服务入门第一篇文章
http://blog.cuicc.com/blog/2015/07/22/microservices/

本系列参考文章
中文地址

https://wuyin.io/archives/page/2/

https://studygolang.com/articles/12060

英文地址（优先推荐）

https://ewanvalentine.io/microservices-in-golang-part-1/
# 笔记

1. protoc.exe .\consignment.proto --go_out=./
这样是错误的，没有写编译插件

2. windows10上makefile的使用
下载软件 http://sourceforge.net/projects/mingw/files/latest/download?source=files

安装，把安装目录的`bin` 目录添加到环境变量。

然后把默认的编译器都安装上即可（其实只要选择编译器是勾选C Compiler 与C++ Compiler）

最后在MinGW的安装目录，打开bin文件夹，将mingw32-make.exe重命名为make.exe

3. 停止当前所有运行中的容器
`docker stop $(docker ps -qa)`

# 学习记录
## dev1

## dev2