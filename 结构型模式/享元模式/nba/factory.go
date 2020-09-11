package nba

import "fmt"

// NBA中的一场比赛由两个球队，主场球队和客场球队，完成比赛，对应着代码就是，一个Match实例会持有2个Team实例。
// 目前，NBA总共由30支球队，按照每个赛季每个球队打82场常规赛算，一个赛季总共会有2460场比赛，对应地，就会有4920个Team实例。
// 但是，NBA的30支球队是固定的，实际上只需30个Team实例就能完整地记录一个赛季的所有比赛信息，剩下的4890个Team实例属于冗余的数据。

// 这种场景下就适合采用享元模式来进行优化，我们把Team设计成多个Match实例之间的享元。
// 享元的获取通过享元工厂来完成，享元工厂teamFactory的定义如下，Client统一使用teamFactory.TeamOf方法来获取球队Team实例。
// 其中，每个球队Team实例只会创建一次，然后添加到球队池中，后续获取都是直接从池中获取，这样就达到了共享的目的。

type teamFactory struct {
	// 球队池，缓存球队实例
	teams map[TeamId]*Team
}

// 根据TeamId创建Team实例，只在TeamOf方法中调用，外部不可见
func createTeam(id TeamId) *Team {
	switch id {
	case Warrior:
		w := &Team{
			Id:   Warrior,
			Name: "荆州勇士",
		}
		curry := &Player{
			Name: "Stephen Curry",
			Team: Warrior,
		}
		thompson := &Player{
			Name: "Klay Thompson",
			Team: Warrior,
		}
		w.Players = append(w.Players, curry, thompson)
		return w
	case Laker:
		l := &Team{
			Id:   Laker,
			Name: "Los Angeles Lakers",
		}
		james := &Player{
			Name: "LeBron James",
			Team: Laker,
		}
		davis := &Player{
			Name: "Anthony Davis",
			Team: Laker,
		}
		l.Players = append(l.Players, james, davis)
		return l
	case Houston:
		h := &Team{
			Id:   Houston,
			Name: "休斯顿火箭",
		}
		yao := &Player{
			Name: "姚明",
			Team: Houston,
		}
		h.Players = append(h.Players, yao)
		return h
	default:
		fmt.Printf("Get an invalid team id %v.\n", id)
		return nil
	}
}

// 根据TeamId获取Team实例，从池中获取，如果池里没有，则创建
func (t *teamFactory) TeamOf(id TeamId) *Team {
	team, ok := t.teams[id]
	if !ok {
		team = &Team{Id: id}
		t.teams[id] = team
		return team
	}
	return team
}

// 享元工厂模式的单例
var factory = &teamFactory{
	teams: make(map[TeamId]*Team),
}

func Factory() *teamFactory {
	return factory
}
