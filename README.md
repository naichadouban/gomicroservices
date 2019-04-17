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
### docker的使用
1. Docker-从入门到实践
https://yeasy.gitbooks.io/docker_practice/content/

## dev3
### yaml文件语法
http://www.ruanyifeng.com/blog/2016/07/yaml.html

### docker-compose的使用
1. 常用命令
`docker-compose build`
`docker-compose up` 加 -d  表示在后台运行
`docker-compose run *`
`docker-compose down` 和run命令相对应
`docker-compose images` 列出compose文件包含的镜像
`docker-compose logs` 列出容器日志
2. 基本参数
-p :表示项目名，不指定的话，默认用所在目录作用项目名。

