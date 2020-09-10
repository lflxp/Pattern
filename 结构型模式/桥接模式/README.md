# Go实现

回到消息处理系统的例子，一个Pipeline对象主要由Input、Filter、Output三类插件组成（3个特征），因为是插件化的系统，不可避免的就要求支持多种Input、Filter、Output的实现，并能够灵活组合（有多个变化的方向）。显然，Pipeline就非常适合使用桥接模式进行设计，实际上我们也这么做了。我们将Input、Filter、Output分别设计成一个抽象的接口，它们按照各自的方向去扩展。Pipeline只依赖的这3个抽象接口，并不感知具体实现的细节。