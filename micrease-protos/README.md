# meshop-protos

protobuf协议项目,这里集中管理所有微服务中的proto文件和生成的代码

### protoc生成
```shell
#生成product中的所有protos,在meshop-protos目录下执行:
#micro
protoc -I ./  --go_out=paths=source_relative:pb/  --micro_out=plugins=micro,paths=source_relative:pb/ product/*.proto
#gateway
protoc -I ./  --go_out=paths=source_relative:gw/  --go-grpc_out=paths=source_relative:gw/   --grpc-gateway_out=logtostderr=true,paths=source_relative:gw/ product/*.proto
```
gopath说明:
* paths=source_relative 表示以.proto文件的目录为相对目录。product/product.proto 生成在

```shell
# import方式生成
protoc -I ./  --go_out=paths=import:pb/  --micro_out=plugins=micro,paths=import:pb/  protos/product/*.proto
```

```shell
# import方式生成
protoc -I ./  --go_out=module=pb  --micro_out=plugins=micro,module=pb  protos/product/*.proto
```

### protoc相关安装
1,安装golang环境  
2,配置go
```shell
#以类linux说明:
export GOROOT=/usr/local/go #设置为go安装的路径
export GOPATH=$HOME/gopath #默认安装包的路径,go get下载的依赖包会存放在此目录中
export GOBIN=$GOPATH/bin #默认安装包的路径,go install安装的可执行命令所在地址
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin #让go install安装的命令生效

#可将上述命令配置在/etc/profile文件中。source /etc/profile使其生效
#PATH也可以配置在/etc/paths文件中
#Mac中开机自动设置环境变量:/etc/zsshc中加入source /etc/profile
#windows中设置环境变量  
#查看go环境
go env
```
3,以go-micro官方安装说明为标准安装
```shell
#protoc,解压后需设置环境变量/etc/paths中加入bin目录
下载地址:https://github.com/protocolbuffers/protobuf/releases 
#protoc-gen-micro
go install go-micro.dev/v4/cmd/protoc-gen-micro@v4
#grpc-gateway
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

网络文档:  
go-micro:https://github.com/asim/go-micro  
protoc-gen-micro:https://github.com/asim/go-micro/tree/master/cmd/protoc-gen-micro

GRPC Gateway:  
参考go-micro example中的说明:https://github.com/go-micro/examples/tree/main/gateway      
grpc-gateway地址:https://github.com/grpc-ecosystem/grpc-gateway  

protoc安装相关:https://blog.csdn.net/junmoxi/article/details/103967504