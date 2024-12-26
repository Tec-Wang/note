# 分布式事务

测试分布式事务框架，以及学习分布式事务专用

# 环境准备
- mysql database1
- mysql database2
- server1
使用grpc进行通信
- server2

启动数据库
mysqld --defaults-file=/cloudide/workspace/mysql/my.cnf --user=$USER --datadir=/cloudide/workspace/mysql/data --socket=/cloudide/workspace/mysql/mysql.sock

链接数据库
mysql  --socket=/cloudide/workspace/mysql/mysql.sock -uroot -p  

# proto文件
https://blog.csdn.net/2301_77868664/article/details/143559151
1. 创建命令
- 普通的go文件，包含了定义的结构体
protoc -I . helloworld.proto --go_out=.
- 服务端需要实现的代码，和客户端创建调用的代码，都在这里面
protoc -I . helloworld.proto --go-grpc_out=.

protoc -I . helloworld.proto --go_out=. --go-grpc_out=.
