package db

// Key-Value数据库接口
type KvDb interface {
	// 存储数据
	// 其中reply为操作结果，存储成功为true，否则为false
	// 当连接数据库失败时返回error，成功则返回nil
	Save(record Record, reply *bool) error
	// 根据key获取value，其中value通过函数参数中指针类型返回
	// 当连接数据库失败时返回error，成功则返回nil
	Get(key string, value *string) error
}

type Record struct {
	Key   string
	Value string
}
