package conngo

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ppkg/glog"
	"xorm.io/xorm"
)

type MySql struct {
	ConnStr string
	engine  *xorm.Engine
}

// 获取一个已存在的Engine
func (m *MySql) GetEngine() *xorm.Engine {
	if m.engine == nil {
		m.engine = m.NewEngine()
	}
	return m.engine
}

// 创建一个新的Engine
func (m *MySql) NewEngine() *xorm.Engine {
	e, err := xorm.NewEngine("mysql", m.ConnStr)

	if err != nil {
		glog.Error(err)
	} else if e != nil {
		location, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			log.Println(err.Error())
		}
		e.SetTZLocation(location)
	}
	return e
}
