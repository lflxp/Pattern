package metric

import (
	"fmt"
	"sync"
)

// 消息输入源统计，设计为单例
type input struct {
	// 存放统计结果，key为Input类型如hello、kafka
	// value为对应Input的消息统计
	metrics map[string]uint64
	// 统计打点时加锁
	mu *sync.Mutex
}

// 给名称为inputName的Input消息计数加1
func (i *input) Inc(inputName string) {
	i.mu.Lock()
	defer i.mu.Unlock()

	if _, ok := i.metrics[inputName]; !ok {
		i.metrics[inputName] = 0
	}

	i.metrics[inputName] = i.metrics[inputName] + 1
}

// 输出当前所有打点的情况
func (i *input) Show() {
	fmt.Printf("Input metric: %v\n", i.metrics)
}

// 单例
var intputInstance = &input{
	metrics: make(map[string]uint64),
	mu:      &sync.Mutex{},
}

func Input() *input {
	return intputInstance
}
