package plugin

import (
	"fmt"
	"reflect"

	msg "github.com/lflxp/123456/创建型模式"
)

// 在输入插件映射关系中加入kafka，用于通过反射创建input对象
func init() {
	inputNames["kafka"] = reflect.TypeOf(KafkaInput{})
}

// 原始struct
type Records struct {
	Items []string
}

type Consumer interface {
	Poll() *Records
}

// 将上面原始struct转换成目标struct KafkaInput
// 重点：转换
// 特殊功能：添加Plugin Func
type KafkaInput struct {
	status   Status
	consumer Consumer
}

func (k *KafkaInput) Receive() *msg.MessageCreater {
	records := k.consumer.Poll()
	if k.status != Started {
		fmt.Println("Kafka input plugin is not running, input nothing.")
		return nil
	}

	return msg.Builder().WithHeaderItem("content", "kafka").WithBodyItem(records.Items[0]).Build()
}

func (k *KafkaInput) Start() {
	k.status = Started
	fmt.Println("KafkaInput plugin started.")
}

func (k *KafkaInput) Stop() {
	k.status = Stopped
	fmt.Println("KafkaInput plugin stopped.")
}

func (k *KafkaInput) Status() Status {
	return k.status
}

// KakkaInput的Init函数实现
func (k *KafkaInput) Init() {
	k.consumer = &MockConsumer{}
}

// 上述代码中的kafka.MockConsumer为我们模式Kafka消费者的一个实现，代码如下
type MockConsumer struct{}

func (m *MockConsumer) Poll() *Records {
	records := &Records{}
	records.Items = append(records.Items, "i am mock consumer.")
	return records
}
