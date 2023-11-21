package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIP   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int // 当前用户的模式
}

func NewClient(serverIP string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		ServerIP:   serverIP,
		ServerPort: serverPort,
		flag:       515,
	}

	// 链接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIP, serverPort))
	if err != nil {
		fmt.Println("net Dial error:", err)
		return nil
	}

	client.conn = conn

	// 返回对象
	return client
}

// 处理server回应的消息，直到显示到标准输出
func (client *Client) DealResponse() {
	// 一旦client.conn有数据，就直接copy到stdout标准输出上，永久阻塞监听
	io.Copy(os.Stdout, client.conn)

	// for {
	// 	buf := make()
	// 	client.conn.Read(buf)
	// 	fmt.Println(buf)
	// }
}

func (client *Client) menu() bool {
	var flag int

	fmt.Println("1. Public chat")
	fmt.Println("2. Private message")
	fmt.Println("3. Update username")
	fmt.Println("0. Exit")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println("Please enter a number within the legal range.")
		return false
	}
}

// 查询在线用户
func (client *Client) SelectUsers() {
	sendMsg := "who"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn Write error:", err)
		return
	}
}

// 私聊模式
func (client *Client) PrivateMessage() {
	var remoteName string
	var chatMsg string

	client.SelectUsers()
	fmt.Println("Please enter the user, exit.")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println("Please enter chat, exit.")
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			// 消息不为空则发送
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn Write error:", err)
					break
				}
			}

			chatMsg = ""
			fmt.Println("Please enter chat, exit.")
			fmt.Scanln(&chatMsg)
		}

		client.SelectUsers()
		fmt.Println("Please enter the user, exit.")
		fmt.Scanln(&remoteName)
	}
}

func (client *Client) PublicChat() {
	// 提示用户输入消息
	var chatMsg string

	fmt.Println("Please enter, exit.")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		// 发给服务器

		// 消息不为空则发送
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn Write error:", err)
				break
			}
		}

		chatMsg = ""
		fmt.Println("Please enter, exit.")
		fmt.Scanln(&chatMsg)
	}

}

func (client *Client) UpdateName() bool {
	fmt.Println("Please enter username:")
	fmt.Scanln(&client.Name)

	sendMsg := "rename|" + client.Name
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn Write error:", err)
		return false
	}

	return true
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {
		}

		// 根据不同的模式处理不同的业务
		switch client.flag {
		case 1:
			// Public Message
			client.PublicChat()
			// break // go的case默认break
		case 2:
			// Private Message
			client.PrivateMessage()
		case 3:
			// Update Username
			client.UpdateName()
		}
	}
}

var serverIP string
var serverPort int

// ./client -ip 127.0.0.1 -port 8888
func init() {
	flag.StringVar(&serverIP, "ip", "127.0.0.1", "Setting the server IP address (default is 127.0.0.1).")
	flag.IntVar(&serverPort, "port", 8888, "Set the server port (default is 8888).")
}

func main() {
	// 命令行解析
	flag.Parse()

	client := NewClient(serverIP, serverPort)
	if client == nil {
		fmt.Println("Connection failure!")
		return
	}

	// 单独开启一个goroutine去处理server的回执消息
	go client.DealResponse()

	fmt.Println("Successfully connected to the server!")

	// 启动客户端业务
	client.Run()
}
