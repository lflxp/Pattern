package nba

// 球队Team的数据结构定义如下
type TeamId uint8

const (
	Warrior TeamId = iota
	Laker
	Houston
)

type Team struct {
	Id      TeamId    // 球队ID
	Name    string    // 球队名称
	Players []*Player // 球队中的球员
}
