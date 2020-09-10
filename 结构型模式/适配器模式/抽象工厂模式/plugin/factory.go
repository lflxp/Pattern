package plugin

import "reflect"

// ============================工厂模式生产具体的类==============================
// 插件抽象工厂接口
type Factory interface {
	Create(conf Config) Plugin
}

// input插件工厂对象，实现Factory接口
type InputFactory struct{}

// 读取配置，通过反射机制进行对象实例化
func (i *InputFactory) Create(conf Config) Plugin {
	t, _ := inputNames[conf.Name]
	p := reflect.New(t).Interface().(Plugin)
	// 返回插件实例前调用Init函数，完成相关初始化方法
	p.Init()
	return p
}

// filter和output插件工厂实现类似
type FilterFactory struct{}

func (f *FilterFactory) Create(conf Config) Plugin {
	t, _ := filterNames[conf.Name]
	return reflect.New(t).Interface().(Plugin)
}

type OutputFactory struct{}

func (o *OutputFactory) Create(conf Config) Plugin {
	t, _ := outputNames[conf.Name]
	return reflect.New(t).Interface().(Plugin)
}
