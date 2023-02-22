# relationmicro
社交服务

# 项目结构
```go
│  handler.go - kitex对外暴露的服务接口
│  main.go    - 主方法, 资源初始化: 注册中心、数据库、线程池等资源
│
├─idl
│      relation.thrift      - idl定义文件
│
├─kitex_gen   - kitex自动生成代码
│  └─relation
│      │  k-consts.go
│      │  k-relation.go
│      │  relation.go
│      │
│      ├─combineservice     - 汇总各模块对外提供
│      │      client.go
│      │      combineservice.go
│      │      invoker.go
│      │      server.go
│      │
│      ├─messageservice     - 消息模块
│      │      client.go
│      │      invoker.go
│      │      messageservice.go
│      │      server.go
│      │
│      └─relationservice    - 关注模块
│              client.go
│              invoker.go
│              relationservice.go
│              server.go
│
├─kitex_server  - 封装其他服务
│      initmicro.go     - 从etcd获取其他服务客户端
│      user_service.go  - 封装usermicro服务接口为本地方法
│      user_service_test.go
│
├─service       - 核心业务逻辑
│      follower_service.go          - 粉丝
│      follower_service_test.go
│      follow_service.go            - 关注
│      follow_service_test.go
│      message_service.go           - 消息
│      message_service_test.go
│      userinfo_service.go
│      user_service.go              - 用户信息
│
└─util
    │  lua_util.go                  - lua脚本相关
    │  redis_constant.go            - redis键生成等固定规则
    │  redis_util.go                - redis调用
    │  redis_util_test.go
    │  str2time_util.go
    │  type_convert_util.go
    │
    └─lua
            follow
            follower
```

# 项目运行
## 运行环境
1. linux`使用windows需配置好远程开发, 因win版的kitex自动生成有bug`
2. go1.19.5
3. redis > 2.6.0 支持 EVALSHA, > 5.0 支持stream
   - 关注用到EVALSHA
   - 消息用到stream 
4. 其他依赖直接拉取最新的即可 `go get & go mod tidy`

## 项目启动需要配置的资源
1. etcd
2. redis
3. usermicro `不启动user服务也能正常跑起来, 但查询列表等功能需要查询user服务，没有user服务会查询失败`
   - 引用usermicro服务`go get github.com/ClubWeGo/usermicro@latest`
上述配置均能再main方法里找到



# 其他
idl 我们使用的kitex+thrift

如需要加入新的rpc接口, 需通过编写thrift文件和kitex命令自动生成相关代码

[kitex官方文档](https://www.cloudwego.io/zh/docs/kitex/)


聚合message的方法到relation的生成代码中的方案
https://www.cloudwego.io/zh/docs/kitex/tutorials/code-gen/combine_service/

使用combine-service 合并多个分块的service
`kitex --combine-service  -service github.com/ClubWeGo/relationmicro -module github.com/ClubWeGo/relationmicro ./idl/relation.thrift`

`go get github.com/cloudwego/kitex@latest && go mod tidy`

注意生成的handler中是CombineServiceImpl
