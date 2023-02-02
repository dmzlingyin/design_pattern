package main

import "fmt"

/*
lighting <---> adapter(transparent) <---> computer
*/

// 客户端接口(需要被适配器实现的接口)
type Computer interface {
	InsertIntoLightningPort()
}

// 客户端代码
type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts lightning connector into computer")
	com.InsertIntoLightningPort()
}

// mac服务(可直接插入lightning)
type Mac struct{}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine")
}

// windows
type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine")
}

type WindowsAdapter struct {
	windowsMachine *Windows
}

func (wa *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB")
	wa.windowsMachine.insertIntoUSBPort()
}

func main() {
	client := &Client{}
	mac := &Mac{}
	client.InsertLightningConnectorIntoComputer(mac)
	fmt.Println("-----------------------------")
	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{windowsMachine}
	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
