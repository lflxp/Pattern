package nba

import (
	"fmt"
	"time"
)

// 比赛结果Match的数据结构定义如下
type Match struct {
	Date         time.Time // 比赛时间
	LocalTeam    *Team     // 主场球队
	VisitorTeam  *Team     // 客场球队
	LocalScore   uint8     // 主场球队得分
	VisitorScore uint8     // 客场球队得分
}

func (m *Match) ShowResult() {
	fmt.Printf("%s VS %s - %d:%d\n", m.LocalTeam.Name, m.VisitorTeam.Name, m.LocalScore, m.VisitorScore)
}
