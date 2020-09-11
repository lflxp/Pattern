package plugin

type Type uint8

const (
	InputType Type = iota
	FilterType
	OutputType
)

// ============================定义接口==============================
// 插件抽象接口定义
type Plugin interface{}

// 输入插件，用于接收消息
type Input interface {
	Plugin
	Receive() string
}

// 过滤插件，用于处理消息
type Filter interface {
	Plugin
	Process(msg string) string
}

// 输入插件，用于发送消息
type Output interface {
	Plugin
	Send(msg string)
}

type Config struct {
	Name       string
	PluginType Type
}
