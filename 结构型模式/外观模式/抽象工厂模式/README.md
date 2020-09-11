# Go实现

外观模式实现起来也很简单，还是考虑前面的消息处理系统。在Pipeline中，每一条消息会依次经过Input->Filter->Output的处理，代码实现起来就是这样：

```go
p := pipeline.Of(config)
message := p.input.Receive()
message = p.filter.Process(message)
p.output.Send(message)
```

但是，对于Pipeline的使用者而言，他可能并不关心消息具体的处理流程，他只需知道消息已经经过Pipeline处理即可。因此，我们需要设计一个简单的对外接口：
```go
package pipeline
...
func (p *Pipeline) Exec() {
 msg := p.input.Receive()
 msg = p.filter.Process(msg)
 p.output.Send(msg)
}
```

这样，使用者只需简单地调用Exec方法，就能完成一次消息的处理，测试代码如下：

```go
package test
...
func TestPipeline(t *testing.T) {
 p := pipeline.Of(pipeline.HelloConfig())
 p.Start()
  // 调用Exec方法完成一次消息的处理
 p.Exec()
}
```

```bash
// 运行结果
=== RUN   TestPipeline
console output plugin started.
upper filter plugin started.
hello input plugin started.
Pipeline started.
Output:
 Header:map[content:text input:hello], Body:[HELLO WORLD]
--- PASS: TestPipeline (0.00s)
PASS
```