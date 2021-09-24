package conngo

import (
	"reflect"
	"testing"

	"github.com/olivere/elastic/v7"
)

var es Es = Es{
	url:      "http://10.1.1.248:9201",
	username: "elastic",
	password: "",
}

func TestEs_GetClient(t *testing.T) {
	tests := []struct {
		name string
		es   *Es
		want *elastic.Client
	}{
		{
			name: "获取一个现有连接",
			es:   &es,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.es.GetClient()
			if got == nil {
				t.Error("got == nil")
			}

			if tt.es.client == nil {
				t.Error("tt.es.client == nil")
			}
			got2 := tt.es.GetClient()
			if !reflect.DeepEqual(got, got2) {
				t.Error("got not equal got2")
			}
		})
	}
}

func TestEs_NewClient(t *testing.T) {
	tests := []struct {
		name string
		es   *Es
		want *elastic.Client
	}{
		{
			name: "创建一个新的连接",
			es:   &es,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.es.NewClient()

			if got == nil {
				t.Error("got == nil")
			}

			if tt.es.client != nil {
				t.Error("tt.es.client != nil")
			}
			got2 := tt.es.NewClient()
			if reflect.DeepEqual(got, got2) {
				t.Error("got equal got2")
			}
		})
	}
}
