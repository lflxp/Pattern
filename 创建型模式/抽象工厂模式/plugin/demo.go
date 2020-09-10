package plugin

import (
	"fmt"
	"reflect"
	"strings"
)

// ============================根据接口实现struct==============================
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
