package beans

import (
	"testing"
)

func TestDefaultBeanFactory_RegisterGetObject(t *testing.T) {
	type args struct {
		name string
		bean interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{name: "bean1", bean: &struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDefaultBeanFactory()
			d.RegisterObject(tt.args.name, tt.args.bean)
			if bean := d.GetObject(tt.args.name); bean != tt.args.bean {
				t.Errorf("GetObject() = %v, want %v", bean, tt.args.bean)
			}
		})
	}
}

func TestNewDefaultBeanFactory(t *testing.T) {
	var tests = []struct {
		name string
	}{
		{
			name: "test1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultBeanFactory(); got == nil || got.beans == nil {
				t.Errorf("NewDefaultBeanFactory() = %v, factory or beans empty", got)
			}
		})
	}
}
