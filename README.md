# 即时通信系统 Instant messaging system

![image-20231119224420622](C:\Users\Yining\AppData\Roaming\Typora\typora-user-images\image-20231119224420622.png)

## 版本迭代

1. 构建基础Server
2. 用户上线功能
3. 用户消息广播机制
4. 用户业务层封装
5. 在线用户查询
6. 修改用户名
7. 超时强踢功能
8. 私聊功能
9. 客户端实现

## 版本一：构建基础Server

- main.go
- server.go
  - server类型
  - 方法
    - func NewServer(ip string, port int) *Server：创建一个server对象
    - func (this *Server) Start()：启动Server服务
    - func (this *Server) Handler(conn net.Conn)：处理链接业务

## 版本二：用户上线功能

![image-20231120021215663](C:\Users\Yining\AppData\Roaming\Typora\typora-user-images\image-20231120021215663.png)

- user.go
  - user类型
  - 方法
    - 创建一个user对象
    - 监听user对应的channel消息
- server.go
  - server类型：新增了OnlineMap和Message属性
  - 在处理客户端上线的Handler创建并添加用户
  - 新增广播消息方法
  - 新增监听广播消息channel方法
  - 用一个goroutine单独监听Message

## 版本三：用户消息广播机制

- server.go
  - 完善handler处理业务方法，启动一个针对当前客户端的读goroutine

## 版本四：用户业务层封装

- user.go
  - user类型新增server关联
  - 新增Online方法
  - 新增Offline方法
  - 新增HandleMessage方法
- server.go
  - 将之前user的业务进行替换

## 版本五：在线用户查询

