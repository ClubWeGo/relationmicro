20230215
~~1、HMGet fields 直接传参异常~~
2、user模块 根据userIds 查询用户已完成 需对接
3、lua - 保证关注 取关的原子性
4、所有的userNames 存在 一个hash结构里太大，想办法在业务层分片一下

<<<<<<< Updated upstream
5、用户name 是不是user服务更新userinfo时去存&redis定时去拉 双重保证比较好
6、kitex




# http层

1、关注操作

2、关注列表

3、粉丝列表

4、好友列表


# rpc层

> rpc层做好参数校验 
> 如判断我是否关注对方 myId和targetId不能是一样的
> userId等字段不能不合理 id > 0


1、关注

2、关注列表 #开发完成、未测试

3、粉丝列表 #开发完成、未测试

4、好友列表（service拿到粉丝列表，然后rpc请求各用户头像url）

5、提供判断是否关注对方的rpc接口给user服务

# service层 

1、关注

2、取关

~~3、 获取关注列表~~

~~4、获取粉丝列表~~

~~5、判断是否关注对方~~



