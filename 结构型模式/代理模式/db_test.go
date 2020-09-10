package 代理模式

import (
	"testing"

	"github.com/lflxp/123456/结构型模式/代理模式/db"
	_ "github.com/lflxp/123456/结构型模式/代理模式/plugin"
	"github.com/lflxp/123456/结构型模式/适配器模式/抽象工厂模式/pipeline"
	"github.com/lflxp/123456/结构型模式/适配器模式/抽象工厂模式/plugin"
)

func TestDbOutput(t *testing.T) {
	db.Start()
	config := pipeline.Config{
		Name: "pipeline3",
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
			Name:       "db",
		},
	}

	p := pipeline.Of(config)
	p.Init()
	p.Start()
	p.Exec()

	// 验证DbOutput存储的正确性
	cli := db.CreateClient()
	var val string
	err := cli.Get("db", &val)
	if err != nil {
		t.Errorf("Get db failed, error: %v\n", err)
	}
	if val != "HELLO WORLD" {
		t.Errorf("expect HELLO WORLD, but actual %s", val)
	}
}
