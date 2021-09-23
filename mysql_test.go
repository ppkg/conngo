package conngo

import (
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var m MySql = MySql{ConnStr: "root:root@(127.0.0.1:3306)/datacenter?charset=utf8"}

func TestMySql_GetEngine(t *testing.T) {
	tests := []struct {
		name string
		m    *MySql
		want *xorm.Engine
	}{
		{
			name: "获取一个现有Engine",
			m:    &m,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.GetEngine()
			if got == nil {
				t.Error("got is nil")
				t.FailNow()
			}

			if tt.m.engine == nil {
				t.Error("tt.m.engine is nil")
				t.FailNow()
			}

			if err := got.Ping(); err != nil {
				t.Error(err)
				t.FailNow()
			}

			if err := got.DB().Ping(); err != nil {
				t.Error(err)
				t.FailNow()
			}
			got2 := tt.m.GetEngine()
			if !reflect.DeepEqual(got, got2) {
				t.Error("got not equal got2")
			}
		})
	}
}

func TestMySql_NewEngine(t *testing.T) {
	tests := []struct {
		name string
		m    *MySql
		want *xorm.Engine
	}{
		{
			name: "获取一个现有Engine",
			m:    &m,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.NewEngine()
			if got == nil {
				t.Error("got is nil")
				t.FailNow()
			}

			if tt.m.engine != nil {
				t.Error("tt.m.engine is not nil")
				t.FailNow()
			}

			if err := got.Ping(); err != nil {
				t.Error(err)
				t.FailNow()
			}

			if err := got.DB().Ping(); err != nil {
				t.Error(err)
				t.FailNow()
			}
			got2 := tt.m.NewEngine()
			if reflect.DeepEqual(got, got2) {
				t.Error("got equal got2")
			}
		})
	}
}
