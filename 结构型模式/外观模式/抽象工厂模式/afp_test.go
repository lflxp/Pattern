package 抽象工厂模式

import (
	"testing"

	pipeline "github.com/lflxp/Pattern/创建型模式/抽象工厂模式/pipeline"
)

func TestPipeline(t *testing.T) {
	// 其中pipeline.DefaultConfig()的配置内容见【抽象工厂模式示例图】
	// 消息处理流程为 HelloInput -> UpperFilter -> ConsoleOutput
	p := pipeline.Of(pipeline.DefaultConfig())
	// 调用Exec方法完成一次消息的处理
	p.Exec()
}
