package conngo

import (
	"bytes"

	"github.com/ppkg/glog"
	"github.com/streadway/amqp"
)

type Mq struct {
	UserName string
	PassWord string
	HostName string
	EndPoint string
	conn     *amqp.Connection
	ch       *amqp.Channel
}

// 获取一个已存在的连接
func (m *Mq) GetConn() *amqp.Connection {
	if m.conn == nil {
		m.conn = m.NewConn()
	}
	return m.conn
}

// 创建一个新的连接
func (m *Mq) NewConn() *amqp.Connection {
	var buf bytes.Buffer

	buf.WriteString("amqp://")
	buf.WriteString(m.UserName)
	buf.WriteString(":")
	buf.WriteString(m.PassWord)

	buf.WriteString("@")
	buf.WriteString(m.HostName)
	buf.WriteString("/")
	buf.WriteString(m.EndPoint)

	url := buf.String()
	if conn, err := amqp.Dial(url); err != nil {
		glog.Error(err)
		return nil
	} else {
		return conn
	}
}

// 获取一个已经存在的 Channel
func (m *Mq) GetChannel(conn *amqp.Connection) *amqp.Channel {
	if m.ch == nil {
		m.ch = m.NewChannel(conn)
	}
	return m.ch
}

// 创建一个新的 Channel
func (m *Mq) NewChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	if err != nil {
		glog.Error(err)
		return nil
	} else {
		return ch
	}
}
