package nba

import (
	"testing"
	"time"
)

func TestNba(t *testing.T) {
	game1 := &Match{
		Date:         time.Date(2020, 1, 10, 9, 30, 0, 0, time.Local),
		LocalTeam:    Factory().TeamOf(Warrior),
		VisitorTeam:  Factory().TeamOf(Laker),
		LocalScore:   102,
		VisitorScore: 99,
	}
	game1.ShowResult()
	game2 := &Match{
		Date:         time.Date(2020, 1, 12, 9, 30, 0, 0, time.Local),
		LocalTeam:    Factory().TeamOf(Laker),
		VisitorTeam:  Factory().TeamOf(Warrior),
		LocalScore:   110,
		VisitorScore: 118,
	}

	game2.ShowResult()
	// 两个Match的同一个球队应该是同一个实例的
	if game1.LocalTeam != game2.VisitorTeam {
		t.Errorf("Warrior team do not use flyweight pattern")
	}
}
