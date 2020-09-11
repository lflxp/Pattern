package plugin

import (
	msg "github.com/lflxp/123456/创建型模式"
	"github.com/lflxp/123456/结构型模式/装饰器模式/抽象工厂模式/metric"
)

type InputMetricDecorator struct {
	input Input
}

func (i *InputMetricDecorator) Receive() *msg.MessageCreater {
	// 调用本地对象的Receive方法
	record := i.input.Receive()
	// 完成统计逻辑
	if inputName, ok := record.Header.Items["content"]; ok {
		metric.Input().Inc(inputName)
	}
	return record
}

func (i *InputMetricDecorator) Start() {
	i.input.Start()
}

func (i *InputMetricDecorator) Stop() {
	i.input.Stop()
}

func (i *InputMetricDecorator) Status() Status {
	return i.input.Status()
}

func (i *InputMetricDecorator) Init() {
	i.input.Init()
}

// 工厂方法，完成装饰器的创建
func CreateInputMetricDecorator(input Input) *InputMetricDecorator {
	return &InputMetricDecorator{input: input}
}
