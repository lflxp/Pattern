package 创建型模式

import (
	"fmt"
	"reflect"
	"strings"
)

// ===================Plugin=================
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

// =======================Pipeline=====================
// 消息管道的定义
type Pipeline struct {
	input  Input
	filter Filter
	output Output
}

// 一个消息的处理流程为 input ->  filter -> output
func (p *Pipeline) Exec() {
	msg := p.input.Receive()
	msg = p.filter.Process(msg)
	p.output.Send(msg)
}

// 接着，我们定义input、filter、output三类插件接口的具体实现：
// input插件名称与类型的映射关系，主要用于通过反射创建input对象
var inputNames = make(map[string]reflect.Type)

// Hello input插件，接收“Hello World”消息
type HelloInput struct{}

func (h *HelloInput) Receive() string {
	return "Hello World"
}

// 初始化input插件映射关系表
func init() {
	inputNames["hello"] = reflect.TypeOf(HelloInput{})
}

// filter插件名称与类型的映射关系，主要用于通过反射创建filter对象
var filterNames = make(map[string]reflect.Type)

// Upper filter插件，将消息全部字母转成大写
type UpperFilter struct{}

func (u *UpperFilter) Process(msg string) string {
	return strings.ToUpper(msg)
}

// 初始化filter插件映射关系表
func init() {
	filterNames["upper"] = reflect.TypeOf(UpperFilter{})
}

// output插件名称与类型的映射关系，主要用于通过反射创建output对象
var outputNames = make(map[string]reflect.Type)

// Console output插件，将消息输出到控制台上
type ConsoleOutput struct{}

func (c *ConsoleOutput) Send(msg string) {
	fmt.Println(msg)
}

// 初始化output插件映射关系表
func init() {
	outputNames["console"] = reflect.TypeOf(ConsoleOutput{})
}

type Config struct {
	Name string
}

// 然后，我们定义插件抽象工厂接口，以及对应插件的工厂实现
// 插件抽象工厂接口
type FactoryAbstract interface {
	Create(conf Config) Plugin
}

// input插件工厂对象，实现FactoryAbstract接口
type InputFactoryAbstract struct{}

// 读取配置，通过反射机制进行对象实例化
func (i *InputFactoryAbstract) Create(conf Config) Plugin {
	t, _ := inputNames[conf.Name]
	return reflect.New(t).Interface().(Plugin)
}

// filter和output插件工厂实现类似
type FilterFactoryAbstract struct{}

func (f *FilterFactoryAbstract) Create(conf Config) Plugin {
	t, _ := filterNames[conf.Name]
	return reflect.New(t).Interface().(Plugin)
}

type OutputFactoryAbstract struct{}

func (o *OutputFactoryAbstract) Create(conf Config) Plugin {
	t, _ := outputNames[conf.Name]
	return reflect.New(t).Interface().(Plugin)
}

// 最后定义pipeline的工厂方法，调用plugin.Factory抽象工厂完成pipelien对象的实例化
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