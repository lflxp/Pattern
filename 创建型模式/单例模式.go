package 创建型模式

import "sync"

// ================饿汉模式===============
// 消息
type Message struct {
	Count int64
}

// 消息池
type messagePool struct {
	pool *sync.Pool
}

// 消息池单例
var msgPool = &messagePool{
	// 如果消息池里没有消息，则新建一个Count值为0的Message实例
	pool: &sync.Pool{New: func() interface{} { return &Message{Count: 0} }},
}

// 访问消息池单例的唯一方法
func Instance() *messagePool {
	return msgPool
}

// 往消息池里添加消息
func (m *messagePool) AddMsg(msg *Message) {
	m.pool.Put(msg)
}

// 从消息池里获取消息
func (m *messagePool) GetMsg() *Message {
	return m.pool.Get().(*Message)
}

// 单例模式的“懒汉模式”实现
var once = &sync.Once{}

// 消息池单例,在首次调用时初始化
var msgPool2 *messagePool

// 全局唯一获取消息池pool的方法
func Instance2() *messagePool {
	once.Do(func() {
		msgPool2 = &messagePool{
			// 如果消息池里没有消息，则新建一个Count值为0的Message实例
			pool: &sync.Pool{New: func() interface{} {
				return &Message{Count: 0}
			}},
		}
	})
	return msgPool2
}
