package plugin

import (
	"fmt"

	. "github.com/lflxp/Pattern/结构型模式/适配器模式/抽象工厂模式/plugin"
)

type LifeCycle struct {
	name   string
	status Status
}

func (l *LifeCycle) Start() {
	l.status = Started
	fmt.Printf("%s plugin started.\n", l.name)
}

func (l *LifeCycle) Stop() {
	l.status = Stopped
	fmt.Printf("%s plugin stopped.\n", l.name)
}

func (l *LifeCycle) Status() Status {
	return l.status
}
