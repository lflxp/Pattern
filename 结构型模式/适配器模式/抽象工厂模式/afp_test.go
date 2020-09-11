package 抽象工厂模式

import (
	"testing"

	"github.com/lflxp/Pattern/结构型模式/适配器模式/抽象工厂模式/pipeline"
	"github.com/lflxp/Pattern/结构型模式/适配器模式/抽象工厂模式/plugin"
)

func TestKafkaInputPipeline(t *testing.T) {
	config := pipeline.Config{
		Name: "pipeline2",
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
	p := pipeline.Of(config)
	p.Start()
	p.Exec()
	p.Stop()
}
