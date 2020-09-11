# Go实现

假设现在需要设计一个系统，用于记录NBA中的球员信息、球队信息以及比赛结果。

球队Team的数据结构定义如下：
```go
package nba
...
type TeamId uint8

const (
 Warrior TeamId = iota
 Laker
)

type Team struct {
 Id      TeamId    // 球队ID
 Name    string    // 球队名称
 Players []*Player // 球队中的球员
}
```
球员Player的数据结构定义如下：
```go
package nba
...
type Player struct {
 Name string // 球员名字
 Team TeamId // 球员所属球队ID
}
```
比赛结果Match的数据结构定义如下：
```go
package nba
...
type Match struct {
 Date         time.Time // 比赛时间
 LocalTeam    *Team     // 主场球队
 VisitorTeam  *Team     // 客场球队
 LocalScore   uint8     // 主场球队得分
 VisitorScore uint8     // 客场球队得分
}

func (m *Match) ShowResult() {
 fmt.Printf("%s VS %s - %d:%d\n", m.LocalTeam.Name, m.VisitorTeam.Name,
  m.LocalScore, m.VisitorScore)
}
```

NBA中的一场比赛由两个球队，主场球队和客场球队，完成比赛，对应着代码就是，一个Match实例会持有2个Team实例。目前，NBA总共由30支球队，按照每个赛季每个球队打82场常规赛算，一个赛季总共会有2460场比赛，对应地，就会有4920个Team实例。但是，NBA的30支球队是固定的，实际上只需30个Team实例就能完整地记录一个赛季的所有比赛信息，剩下的4890个Team实例属于冗余的数据。

这种场景下就适合采用享元模式来进行优化，我们把Team设计成多个Match实例之间的享元。享元的获取通过享元工厂来完成，享元工厂teamFactory的定义如下，Client统一使用teamFactory.TeamOf方法来获取球队Team实例。其中，每个球队Team实例只会创建一次，然后添加到球队池中，后续获取都是直接从池中获取，这样就达到了共享的目的。