# 1.完成文档名修改
# 2.完成所有单元测试

test: create struct
	@echo 开始测试
	pwd

create:
	@echo 创建型模式
	cd 创建型模式 && go test -v
	cd 创建型模式/抽象工厂模式 && go test -v

struct:
	@echo 结构型模式
	cd 结构型模式/代理模式 && go test -v
	cd 结构型模式/适配器模式/抽象工厂模式 && go test -v
	cd 结构型模式/外观模式/抽象工厂模式 && go test -v
	cd 结构型模式/享元模式/nba && go test -v
	cd 结构型模式/装饰器模式/抽象工厂模式 && go test -v
	cd 结构型模式/组合模式/抽象工厂模式 && go test -v

prepare:
	@echo 修改中文路径缺失问题
	cd .. && mv - 123456
