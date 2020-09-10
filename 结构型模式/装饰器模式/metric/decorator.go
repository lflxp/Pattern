package metric

import (
	msg "github.com/lflxp/123456/创建型模式"
)

type InputMetricDecorator struct {
	input input
}

func (i *InputMetricDecorator) Receive() *msg.MessageCreater {

}
