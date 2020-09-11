package plugin

import (
	msg "github.com/lflxp/Pattern/创建型模式"
)

type Type uint8

const (
	InputType Type = iota
	FilterType
	OutputType
)

// 添加插件运行状态
type Status uint8

const (
	Stopped Status = iota
	Started
)

// ============================定义接口==============================
// Input,Filter,Output三类插件接口的定义跟上一篇文章类似
// 这里使用MessageCreater结构体替代里原来的string，使得语义更清晰
// 插件抽象接口定义
type Plugin interface {
	// 启动插件
	Start()
	// 停止插件
	Stop()
	// 返回插件当前的运行状态
	Status() Status
	// 新增初始化方法，在插件工厂返回实例前调用
	Init()
}

// 输入插件，用于接收消息
type Input interface {
	Plugin
	Receive() *msg.MessageCreater
}

// 过滤插件，用于处理消息
type Filter interface {
	Plugin
	Process(msg *msg.MessageCreater) *msg.MessageCreater
}

// 输入插件，用于发送消息
type Output interface {
	Plugin
	Send(msg *msg.MessageCreater)
}

type Config struct {
	Name       string
	PluginType Type
}
