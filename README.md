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

3. 停止当前所有运行中的容器、删除镜像
`docker stop $(docker ps -aq)`
`docker rmi $(docker images -q)`
`docker container rm *` 删除一个处于终止状态的容器
`docker container prune` 清除所有处于终止状态的容器 



# 学习记录
## dev1
这一节基本上还是grpc的内容

1. 微服务的介绍
2. grpc的使用介绍
3. protobuf的使用

### consignment微服务
1. 创建protobuf服务描述文件consignment.proto
2. 用protoc把consignment.proto编译成go版本
3. 创建grpc服务端

关键代码
```go
listener, err := net.Listen("tcp", PORT)
server := grpc.NewServer()
pb.RegisterShippingServiceServer(server, &service{repo})
server.Serve(listener)
```
4. 创建grpc客户端

关键代码
```go
conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
client := pb.NewShippingServiceClient(conn)
// 调用
resp, err := client.CreateConsignment(context.Background(), consignment)
```


## dev2
### docker
1. Docker-从入门到实践
https://yeasy.gitbooks.io/docker_practice/content/

docker只是容器技术的一种实现，比较流行的一种。不要把容器技术和docker混淆了。
### 把consignment微服务容器化
创建Dockerfile文件
```docker
# 若运行环境是 Linux 则需把 alpine 换成 debian
# 使用最新版 alpine 作为基础镜像
FROM alpine:latest

# 在容器的根目录下创建 app 目录
RUN mkdir /app

# 将工作目录切换到 /app 下
WORKDIR /app

# 将微服务的服务端运行文件拷贝到 /app 下
ADD consignment-service /app/consignment-service

# 运行服务端
CMD ["./consignment-service"]
```
为了使用方便，我们创建Makefile文件
```docker
build:
	# 编译proto文件（当前应该是grpc，micro是后面加上go-micro之后的）
	protoc -I. --go_out=plugins=micro:. proto/consignment/consignment.proto
	# 编译服务代码
	GOOS=linux GOARCH=amd64 go build
	# 创建docker镜像
	docker build -t shippy-service-consignment .
run:
	# 启动docker容器
	docker run -p 8010:8010 -e MICRO_SERVER_ADDRESS=:8010 shippy-service-consignment

```
### go-micro
> 为什么需要go-micro，我们用grpc不就挺好吗？
1. 管理麻烦:我们在consignment/cli.go文件中，就是调用服务时，需要手动指定服务的ip和端口。当我们部署很多服务后，一个服务的IP或者端口发生变化，其他服务就都不能调用了。
2. 服务发现：为了解决各个微服务间的调用问题，就需要服务发现。服务发现就相当于是一个注册中心，各个微服务启动时都到这里来记录自己的IP和端口，服务注销时，也把记录的IP和端口注销掉。go-micro刚好实现了这一点（不止这一点），我们就用go-micro来改造我们的服务。

### 用go-micro改造consignment服务端

1. 首先consignment.proto 用protoc编译时，plugins插件改成了micro

用插件 micro 插件编译后，之前.proto 定义的方法中的返回值现在移到了请求参数中。
2. 服务端代码也需要改

主要代码
```go
server := micro.NewService(
		micro.Name("shippy.service.consignmen"),
		micro.Version("latest"),
    )
pb.RegisterShippingServiceHandler(server.Server(), &service{repo})
server.Run()
```
3. 监听地址没有写死

查看代码发现，我们再go-micro服务启动的过程中，自始至终都没有出现服务监听的ip和端口。why ？
> go-mirco 会自动使用系统或命令行中变量 MICRO_SERVER_ADDRESS 的地址

所以我们再Makefile中docker启动命令中，加了环境变量`-e MICRO_SERVER_ADDRESS=:8010 `

go-micro默认是使用mdns作为服务发现的，我们可以使用环境变量 `MICRO_REGISTRY=mdns` 来设置。在生产环境中，可以使用etcd来代替。

### 用go-micro改造consignment客户端
主要代码

```go
service := micro.NewService(micro.Name("shippy.consignment.cli"))
service.Init()
// 这个名字 shippy.service.consignment 一定要写对啊
client := pb.NewShippingServiceClient("shippy.service.consignment", service.Client())
// 调用
res, err := client.CreateConsignment(context.Background(), consignment)
```

### 用go-micro创建VesselService微服务


## dev3
### yaml文件语法
http://www.ruanyifeng.com/blog/2016/07/yaml.html

### docker-compose的使用
之前，我们启动每一个微服务都需要有一个Makefile文件，如果微服务很多的话，管理起来很不方便。我们就需要一个统一管理:docker-compoose

为了使用docker-compose，我们几个服务的同级目录下创建了 `docker-compose.yaml` 文件。
1. 常用命令

```go
docker-compose build
docker-compose up //加 -d  表示在后台运行
docker-compose run *
docker-compose down //和run命令相对应
docker-compose images //列出compose文件包含的镜像
docker-compose logs //列出容器日志
```

2. 基本参数

-p :表示项目名，不指定的话，默认用所在目录作用项目名。

### mgo

最好的驱动 `go get gopkg.in/mgo.v2`

简单的使用教程： https://www.itfanr.cc/2017/06/28/golang-connect-to-mongodb/

条件查询的示例：https://www.jianshu.com/p/b63e5cfa4ce5

### consignment服务重构

之前一个服务的代码，全部都写在main.go中。仓库repository，服务service逻辑，go-micro服务都在一起显得有点混乱。

现在我们加了mongoDb，想要数据持久化。

所以我们把代码拆分成几个文件：
```go
datastore.go // 创建与 MongoDB 的主回话
repository.go // 实现与数据库进行的操作
handler.go // 实现微服务的服务端，处理业务骆
main.go // 注册 并启动服务

```
### vessel-service重构

重构逻辑与consignment服务相同

### 在docker-compose.yaml中引入MongoDB

```docker
services:
  ...
  datastore:
    image: mongo
    ports:
      - 27017:27017
```
Note:mongoDB 的docker服务是`datastore`。这是用到了docker内置DNS机制。

> 我们的consignment服务，vessel服务中，需要使用MongoDB是，就可以直接这样设置`DB_HOST: "datastore:27017"`

### user-service

#### 首先需要引入docker service：postgresql
```docker
  ...
  database:
    image: postgres
    ports:
      - 5432:5432
```

Note:这个docker 服务名：database,就表示我们在其他地方使用此postgresql，就可以用`database:5432`表示了

#### UserService服务
1. 定义proto文件
2. 实现业务处理逻辑 handler.go
3. 实现数据库的交互 repository.go

这里我们使用了gorm：https://jasperxu.github.io/gorm-zh
4. 