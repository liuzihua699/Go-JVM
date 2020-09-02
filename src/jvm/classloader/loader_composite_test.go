package classloader

import (
	"reflect"
	"testing"
)

func TestComClassLoader_LoadClass(t *testing.T) {
	type fields struct {
		paths   string
		listDir []string
		loaders []ClassLoader
	}
	type args struct {
		className string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     []byte
		want1    ClassLoader
		wantErr  bool
		pathList string
	}{
		// TODO: Add test cases.
		{
			name:     "1",
			args:     args{className: "com.zihua.HelloWorld"},
			pathList: "a;b;c;T:\\\\jvm-test;d",
		},
		{
			name:     "2",
			args:     args{className: "HelloWorld"},
			pathList: "a;b;c;T:\\\\jvm-test;d",
		},
		{
			name:     "3",
			args:     args{className: "java.lang.String"},
			pathList: "a;b;c;T:\\\\jvm-test;d",
		},
		{
			name:     "4",
			args:     args{className: "java.lang.String"},
			pathList: "a;b;c;T:\\\\jvm-test;d;C:\\Program Files\\Java\\jdk1.8.0_161\\jre\\lib\\rt.jar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//c := ComClassLoader{
			//	paths:   tt.fields.paths,
			//	listDir: tt.fields.listDir,
			//	loaders: tt.fields.loaders,
			//}
			c := CreateCompositeLoader(tt.pathList)
			got, got1, err := c.LoadClass(tt.args.className)
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

func TestComClassLoader_ToString(t *testing.T) {
	type fields struct {
		paths   string
		listDir []string
		loaders []ClassLoader
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "1",
			want: "ComClassLoader",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ComClassLoader{
				paths:   tt.fields.paths,
				listDir: tt.fields.listDir,
				loaders: tt.fields.loaders,
			}
			if got := c.ToString(); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateCompositeLoader(t *testing.T) {
	type args struct {
		pathList string
	}
	tests := []struct {
		name string
		args args
		want *ComClassLoader
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{pathList: "a;b;c;T:\\\\jvm-test;d"},
		},
		{
			name: "2",
			args: args{pathList: "a;b;c;d"},
		},
		{
			name: "3",
			args: args{pathList: "a;b;c;T:\\\\jvm-test;d;C:\\Program Files\\Java\\jdk1.8.0_161\\jre\\lib\\rt.jar"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateCompositeLoader(tt.args.pathList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCompositeLoader() = %v, want %v", got, tt.want)
			}
		})
	}
}
