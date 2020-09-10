# afp

`Abstract Factory Pattern`

# 简述

在工厂方法模式中，我们通过一个工厂对象来创建一个产品族，具体创建哪个产品，则通过swtich-case的方式去判断。这也意味着该产品组上，每新增一类产品对象，都必须修改原来工厂对象的代码；而且随着产品的不断增多，工厂对象的职责也越来越重，违反了单一职责原则。

抽象工厂模式通过给工厂类新增一个抽象层解决了该问题，如上图所示，FactoryA和FactoryB都实现·抽象工厂接口，分别用于创建ProductA和ProductB。如果后续新增了ProductC，只需新增一个FactoryC即可，无需修改原有的代码；因为每个工厂只负责创建一个产品，因此也遵循了单一职责原则。

# Go实现

考虑需要如下一个`插件架构风格`的消息处理系统，pipeline是消息处理的管道，其中包含了input、filter和output三个插件。我们需要实现根据配置来创建pipeline ，`加载插件过程`的实现非常`适合使用工厂模式`，其中`input、filter和output三类插件`的创建`使用抽象工厂模式`，而`pipeline的创建`则使用`工厂方法模式`。

# 访问地址

> [使用Go实现GoF的23种设计模式（一）](https://blog.csdn.net/ruanrunxue/article/details/107903077)