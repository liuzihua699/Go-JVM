package classloader

import (
	"reflect"
	"testing"
)

func TestCreateWildcardLoader(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *WildcardClassLoader
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{path: "C:\\Program Files\\Java\\jdk1.8.0_161\\jre\\lib\\*"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateWildcardLoader(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateWildcardLoader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWildcardClassLoader_LoadClass(t *testing.T) {
	type fields struct {
		basePath string
	}
	type args struct {
		className string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		want1   ClassLoader
		wantErr bool
		path    string
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{className: "java.lang.Class"},
			path: "C:\\Program Files\\Java\\jdk1.8.0_161\\jre\\lib\\*",
		},
		{
			name: "2",
			args: args{className: "java.lang.Class"},
			path: "T:\\jvm-test\\*",
		},
		{
			name: "3",
			args: args{className: "com.zihua.HelloWorld"},
			path: "T:\\jvm-test\\*",
		},
		{
			name: "4",
			args: args{className: "java.lang.Class"},
			path: "T:\\jvm-test\\*",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//w := WildcardClassLoader{
			//	basePath: tt.fields.basePath,
			//}
			w := CreateWildcardLoader(tt.path)
			got, got1, err := w.LoadClass(tt.args.className)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadClass() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadClass() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LoadClass() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestWildcardClassLoader_ToString(t *testing.T) {
	type fields struct {
		basePath string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := WildcardClassLoader{
				basePath: tt.fields.basePath,
			}
			if got := w.ToString(); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
