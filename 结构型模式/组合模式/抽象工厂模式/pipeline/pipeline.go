package pipeline

import (
	"fmt"

	plugin "github.com/lflxp/Pattern/结构型模式/组合模式/抽象工厂模式/plugin"
)

// Config
type Config struct {
	Input  plugin.Config
	Filter plugin.Config
	Output plugin.Config
}

func DefaultConfig() Config {
	return Config{
		Input: plugin.Config{
			PluginType: plugin.InputType,
			Name:       "hello",
		},
		Filter: plugin.Config{
			PluginType: plugin.FilterType,
			Name:       "upper",
		},
		Output: plugin.Config{
			PluginType: plugin.OutputType,
			Name:       "console",
		},
	}
}

// 对于插件化的消息处理系统而言，一切皆是插件，因此我们将Pipeine也设计成一个插件，实现Plugin接口
// 消息管道的定义
type Pipeline struct {
	status plugin.Status
	input  plugin.Input
	filter plugin.Filter
	output plugin.Output
}

// 一个消息的处理流程为 input ->  filter -> output
func (p *Pipeline) Exec() {
	msg := p.input.Receive()
	msg = p.filter.Process(msg)
	p.output.Send(msg)
}

// 启动的顺序 output -> filter -> input
func (p *Pipeline) Start() {
	p.output.Start()
	p.filter.Start()
	p.input.Start()
	p.status = plugin.Started
	fmt.Println("Pipeline started.")
}

// 停止的顺序 input -> filter -> output
func (p *Pipeline) Stop() {
	p.input.Stop()
	p.filter.Stop()
	p.output.Stop()
	p.status = plugin.Stopped
	fmt.Println("Pipeline stopped.")
}

func (p *Pipeline) Status() plugin.Status {
	return p.status
}

// 最后定义pipeline的工厂方法，调用plugin.Factory抽象工厂完成pipelien对象的实例化：
// 保存用于创建Plugin的工厂实例，其中map的key为插件类型，value为抽象工厂接口
var pluginFactories = make(map[plugin.Type]plugin.Factory)

// 根据plugin.Type返回对应Plugin类型的工厂实例
func factoryOf(t plugin.Type) plugin.Factory {
	factory, _ := pluginFactories[t]
	return factory
}

// pipeline工厂方法，根据配置创建一个Pipeline实例
func Of(conf Config) *Pipeline {
	p := &Pipeline{}
	p.input = factoryOf(plugin.InputType).Create(conf.Input).(plugin.Input)
	p.filter = factoryOf(plugin.FilterType).Create(conf.Filter).(plugin.Filter)
	p.output = factoryOf(plugin.OutputType).Create(conf.Output).(plugin.Output)
	return p
}

// 初始化插件工厂对象
func init() {
	pluginFactories[plugin.InputType] = &plugin.InputFactory{}
	pluginFactories[plugin.FilterType] = &plugin.FilterFactory{}
	pluginFactories[plugin.OutputType] = &plugin.OutputFactory{}
}
