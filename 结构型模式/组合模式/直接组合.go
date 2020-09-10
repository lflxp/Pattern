package 结构型模式

// 一个典型的实现如《使用Go实现GoF的23种设计模式（一）》中所举的例子，一个Message结构体，由Header和Body所组成。那么Message就是一个整体，而Header和Body则为消息的组成部分。

type A struct{}
type B struct{}

type Message struct {
	a A
	b B
}
