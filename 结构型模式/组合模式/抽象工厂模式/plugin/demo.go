package plugin

import (
	"fmt"
	"reflect"
	"strings"

	msg "github.com/lflxp/123456/创建型模式"
)

// ============================根据接口实现struct==============================
// 接着，我们定义input、filter、output三类插件接口的具体实现：
// input插件名称与类型的映射关系，主要用于通过反射创建input对象
var inputNames = make(map[string]reflect.Type)

// Hello input插件，接收“Hello World”消息
type HelloInput struct {
	status Status
}

func (h *HelloInput) Receive() *msg.MessageCreater {
	// 如果插件未启动，则返回nil
	if h.status != Started {
		fmt.Println("Hello input plugin is not running,input nothing.")
		return nil
	}
	return msg.Builder().
		WithHeaderItem("content", "text").
		WithBodyItem("Hello World").
		Build()
}

func (h *HelloInput) Start() {
	h.status = Started
	fmt.Println("Hello input plugin started.")
}

func (h *HelloInput) Stop() {
	h.status = Stopped
	fmt.Println("Hello input plugin stopped.")
}

func (h *HelloInput) Status() Status {
	return h.status
}

// 初始化input插件映射关系表
func init() {
	inputNames["hello"] = reflect.TypeOf(HelloInput{})
}

// filter插件名称与类型的映射关系，主要用于通过反射创建filter对象
var filterNames = make(map[string]reflect.Type)

// Upper filter插件，将消息全部字母转成大写
type UpperFilter struct {
	status Status
}

func (u *UpperFilter) Process(msgs *msg.MessageCreater) *msg.MessageCreater {
	if u.status != Started {
		fmt.Println("Upper filter plugin is not running ,filter nothing.")
		return msgs
	}

	for i, val := range msgs.Body.Items {
		msgs.Body.Items[i] = strings.ToUpper(val)
	}
	return msgs
}

func (u *UpperFilter) Start() {
	u.status = Started
	fmt.Println("Upper filter plugin started.")
}

func (u *UpperFilter) Stop() {
	u.status = Stopped
	fmt.Println("Upper filter plugin stopped.")
}

func (u *UpperFilter) Status() Status {
	return u.status
}

// 初始化filter插件映射关系表
func init() {
	filterNames["upper"] = reflect.TypeOf(UpperFilter{})
}

// output插件名称与类型的映射关系，主要用于通过反射创建output对象
var outputNames = make(map[string]reflect.Type)

// Console output插件，将消息输出到控制台上
type ConsoleOutput struct {
	status Status
}

func (c *ConsoleOutput) Send(msgs *msg.MessageCreater) {
	if c.status != Started {
		fmt.Println("Console output is not running, output nothing.")
		return
	}
	fmt.Printf("Output:\n\tHeader: %+v, Body: %+v\n", msgs.Header.Items, msgs.Body.Items)
}

func (c *ConsoleOutput) Start() {
	c.status = Started
	fmt.Println("Console output plugin started.")
}

func (c *ConsoleOutput) Stop() {
	c.status = Stopped
	fmt.Println("Console output plugin stopped.")
}

func (c *ConsoleOutput) Status() Status {
	return c.status
}

// 初始化output插件映射关系表
func init() {
	outputNames["console"] = reflect.TypeOf(ConsoleOutput{})
}
