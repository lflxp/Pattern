package 创建型模式

import "testing"

func TestFactoryCreate(t *testing.T) {
	factory := Factory{}
	e := factory.Create(Start)
	if e.EventType() != Start {
		t.Errorf("expect Start, but actual %v.", e.EventType())
	}

	e = factory.Create(End)
	if e.EventType() != End {
		t.Errorf("expect End, but actual %v.", e.EventType())
	}
}
