package plugin

import (
	"fmt"

	msg "github.com/lflxp/Pattern/创建型模式"
	"github.com/lflxp/Pattern/结构型模式/代理模式/db"
	"github.com/lflxp/Pattern/结构型模式/适配器模式/抽象工厂模式/plugin"
)

// 在output中添加dboutput，用于通过反射创建dboutput对象
func init() {
	plugin.AddOutputNames("db", DbOutput{})
}

type DbOutput struct {
	LifeCycle
	// 操作数据库的远程代理
	proxy db.KvDb
}

func (d *DbOutput) Send(msgs *msg.MessageCreater) {
	if d.status != plugin.Started {
		fmt.Printf("%s is not running, output nothing.\n", d.name)
		return
	}

	record := db.Record{
		Key:   "db",
		Value: msgs.Body.Items[0],
	}

	reply := false
	err := d.proxy.Save(record, &reply)
	if err != nil || !reply {
		fmt.Println("Save msg to db server failed.")
	}
}

func (d *DbOutput) Init() {
	fmt.Println("DbOutput Init")
	d.proxy = db.CreateClient()
	d.name = "db output"
}
