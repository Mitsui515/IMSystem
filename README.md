# Instant Messaging System

![Overall Architecture](./Overall_Architecture.png)

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

- user.go
  - 提供SendMsg向对象客户端发送消息API
  - 在HandleMessage()方法中，加上对"who"指令的处理，返回在线用户信息

## 版本六：修改用户名

消息格式"rename|Bob"

- user.go
  - 在HandleMessage()方法中，加上对"rename|Bob"指令的处理，返回在线用户信息

## 版本七：超时强踢功能

用户的任意消息表示用户为活跃，长时间不发消息认为超时，就要强制关闭用户连接

- server.go
  - 在用户的Handler() goroutine中，添加用户活跃channel，一旦有消息，就向该channel发送数据
  - 在用户的Handler() goroutine中，添加定时器功能，超时则强踢

## 版本八：私聊功能

消息格式"to|username|message"

- user.go
  - 在HandleMessage()方法中，加上对"to|username|message"指令的处理，返回在线用户信息

## 版本九：客户端实现

- client.go
  - 客户端类型定义与连接
    - client类型
    - 创建client对象，同时连接服务器
    - 主业务创建client对象
  - 解析命令行
    - init函数初始化命令行参数
    - main函数解析命令行
  - 菜单显示
    - Client新增flag属性
    - 新增menu()方法，获取用户输入的新模式
    - 新增Run()主业务循环
    - main中调用client.Run()
  - 更新用户名
    - 新增UpdateName()更新用户名
    - 加入到Run业务分支中
    - 添加处理server回执消息方法DealResponse()
    - 开启一个goroutine，去承载DealResponse()
  - 公聊模式
    - 新增PublicChat()公聊模式业务
    - 加入Run的分支中
  - 私聊模式
    - 查询当前都有哪些用户在线
    - 提示用户选择一个用户进入私聊

