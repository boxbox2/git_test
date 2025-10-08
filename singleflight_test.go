package singleflight

import (
	"reflect"
	"sync"
	"testing"
)

func TestNewGroup(t *testing.T) {
	tests := []struct {
		name string
		want Group
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGroup(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_group_Do(t *testing.T) {
	type fields struct {
		mu sync.Mutex
		m  map[string]*call
	}
	type args struct {
		key string
		fn  func() (interface{}, error)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := group{
				mu: tt.fields.mu,
				m:  tt.fields.m,
			}
			got, err := g.Do(tt.args.key, tt.args.fn)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_group_DoChan(t *testing.T) {
	type fields struct {
		mu sync.Mutex
		m  map[string]*call
	}
	type args struct {
		key string
		fn  func() (interface{}, error)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   <-chan Result
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := group{
				mu: tt.fields.mu,
				m:  tt.fields.m,
			}
			if got := g.DoChan(tt.args.key, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoChan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_group_Forget(t *testing.T) {
	type fields struct {
		mu sync.Mutex
		m  map[string]*call
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := group{
				mu: tt.fields.mu,
				m:  tt.fields.m,
			}
			g.Forget(tt.args.key)
		})
	}
}

func Test_group_cleanup(t *testing.T) {
	type fields struct {
		mu sync.Mutex
		m  map[string]*call
	}
	type args struct {
		key string
		c   *call
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &group{
				mu: tt.fields.mu,
				m:  tt.fields.m,
			}
			g.cleanup(tt.args.key, tt.args.c)
		})
	}
}

func Test_group_safeCall(t *testing.T) {
	type fields struct {
		mu sync.Mutex
		m  map[string]*call
	}
	type args struct {
		fn func() (interface{}, error)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantVal interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &group{
				mu: tt.fields.mu,
				m:  tt.fields.m,
			}
			gotVal, err := g.safeCall(tt.args.fn)
			if (err != nil) != tt.wantErr {
				t.Errorf("safeCall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("safeCall() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
		})
	}
}

func Test_toError(t *testing.T) {
	type args struct {
		r interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := toError(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("toError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
