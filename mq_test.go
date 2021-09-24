package conngo

import (
	"reflect"
	"testing"

	"github.com/streadway/amqp"
)

var mq Mq = Mq{
	username: "admin",
	password: "admin",
	hostname: "10.1.1.248:5672",
	endpoint: "",
}

func TestMq_GetConn(t *testing.T) {
	tests := []struct {
		name string
		m    *Mq
		want *amqp.Connection
	}{
		{
			name: "获取一个现有的连接",
			m:    &mq,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.GetConn()
			if got == nil {
				t.Error("got = nil")
				t.FailNow()
			}

			if tt.m.conn == nil {
				t.Error("tt.m.conn = nil")
				t.FailNow()
			}

			ch := tt.m.GetChannel(got)
			if ch == nil {
				t.Error("ch = nil")
				t.FailNow()
			}

			if tt.m.ch == nil {
				t.Error("tt.m.ch = nil")
				t.FailNow()
			}
			got2 := tt.m.GetConn()
			if !reflect.DeepEqual(got, got2) {
				t.Error("got not equal got2")
			}
		})
	}
}

func TestMq_NewConn(t *testing.T) {
	tests := []struct {
		name string
		m    *Mq
		want *amqp.Connection
	}{
		{
			name: "创建一个新的连接",
			m:    &mq,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.NewConn()
			if got == nil {
				t.Error("got is nil")
				t.FailNow()
			}

			if tt.m.conn != nil {
				t.Error("tt.m.conn is not nil")
				t.FailNow()
			}

			ch := tt.m.NewChannel(got)
			if ch == nil {
				t.Error("ch is nil")
				t.FailNow()
			}

			if tt.m.ch != nil {
				t.Error("tt.m.ch is not nil")
				t.FailNow()
			}
			got2 := tt.m.NewConn()
			if reflect.DeepEqual(got, got2) {
				t.Error("got equal got2")
			}
		})
	}
}
