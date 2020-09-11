package 抽象工厂模式

import (
	"testing"

	"github.com/lflxp/123456/结构型模式/装饰器模式/抽象工厂模式/metric"
	"github.com/lflxp/123456/结构型模式/装饰器模式/抽象工厂模式/pipeline"
	"github.com/lflxp/123456/结构型模式/装饰器模式/抽象工厂模式/plugin"
)

func TestKafkaInputPipeline(t *testing.T) {
	kafkaconfig := pipeline.Config{
		Name: "pipeline5",
		Input: plugin.Config{
			PluginType: plugin.InputType,
			Name:       "kafka",
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
	helloconfig := pipeline.Config{
		Name: "pipeline6",
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

	p1 := pipeline.Of(helloconfig)
	p2 := pipeline.Of(kafkaconfig)

	p1.Start()
	p2.Start()
	p1.Exec()
	p2.Exec()
	p1.Exec()
	p1.Stop()
	p2.Stop()

	metric.Input().Show()
}
