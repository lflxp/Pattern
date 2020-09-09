package 创建型模式

// ================接口类定义==================
type Type uint8

// 事件类型定义
const (
	Start Type = iota
	End
)

// 事件抽象接口
type Event interface {
	EventType() Type
	Content() string
}

// 开始事件，实现了Event接口
type StartEvent struct {
	content string
}

func (s *StartEvent) EventType() Type {
	return Start
}

func (s *StartEvent) Content() string {
	return s.content
}

// 结束事件，实现了Event接口
type EndEvent struct {
	content string
}

func (s *EndEvent) EventType() Type {
	return End
}

func (s *EndEvent) Content() string {
	return s.content
}

// ================类实现==================
// 事件工厂对象
type Factory struct{}

// 根据事件类型创建具体事件
func (e *Factory) Create(etype Type) Event {
	switch etype {
	case Start:
		return &StartEvent{
			content: "this is start event",
		}
	case End:
		return &EndEvent{
			content: "this is end event",
		}
	default:
		return nil
	}
}

// 按照第二种实现方式，分别给Start和End类型的Event单独提供一个工厂方法
// Start类型Event的工厂方法
func OfStart() Event {
	return &StartEvent{
		content: "this is start event",
	}
}

// End类型Event的工厂方法
func OfEnd() Event {
	return &EndEvent{
		content: "this is end event",
	}
}
