package conngo

import (
	"reflect"
	"testing"

	"github.com/go-redis/redis"
)

var r Redis = Redis{addr: "10.1.1.245:6979", password: "", db: 0}

func TestRedis_GetClient(t *testing.T) {
	tests := []struct {
		name string
		r    *Redis
		want *redis.Client
	}{
		{
			name: "获取一个现有连接",
			r:    &r,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.r.GetClient()
			if got == nil {
				t.Error("got is nil")
				t.FailNow()
			}

			if tt.r.client == nil {
				t.Error("tt.r.client is nil")
				t.FailNow()
			}

			if _, err := got.Ping().Result(); err != nil {
				t.Error(err)
				t.FailNow()
			}

			got2 := tt.r.GetClient()
			if !reflect.DeepEqual(got, got2) {
				t.Error("got not equal got2")
			}

		})
	}
}

func TestRedis_NewClient(t *testing.T) {
	tests := []struct {
		name string
		r    *Redis
		want *redis.Client
	}{
		{
			name: "获取一个现有连接",
			r:    &r,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.r.NewClient()
			if got == nil {
				t.Error("got is nil")
				t.FailNow()
			}

			if tt.r.client != nil {
				t.Error("tt.r.client is not nil")
				t.FailNow()
			}

			if _, err := got.Ping().Result(); err != nil {
				t.Error(err)
				t.FailNow()
			}

			got2 := tt.r.NewClient()
			if reflect.DeepEqual(got, got2) {
				t.Error("got equal got2")
				t.FailNow()
			}

		})
	}
}
