// 还是以建造者模式一节中的Message作为例子，现在设计一个Prototype抽象接口：

package 创建型模式

// 原型复制抽象接口
type Prototype interface {
	Clone() Prototype
}

func (m *MessageCreater) Clone() Prototype {
	msg := m
	return msg
}
