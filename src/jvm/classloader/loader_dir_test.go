package classloader

import (
	"reflect"
	"testing"
)

func TestCreateDirLoader(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *DirClassLoader
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{"T:\\"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateDirLoader(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateDirLoader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirClassLoader_LoadClass(t *testing.T) {
	type fields struct {
		absDir string
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
	}{
		// TODO: Add test cases.
		{
			name:   "1",
			fields: fields{absDir: "T:\\jvm-test"},
			args:   args{className: "HelloWorld.class"},
		},
		{
			name:   "2",
			fields: fields{absDir: "T:\\jvm-test"},
			args:   args{className: "HelloWorld"},
		},
		{
			name:   "3",
			fields: fields{absDir: "T:\\jvm-test"},
			args:   args{className: "com.zihua.HelloWorld"},
		},
		{
			name:   "4",
			fields: fields{absDir: "T:\\jvm-test"},
			args:   args{className: "HelloWorld.txt"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DirClassLoader{
				absDir: tt.fields.absDir,
			}
			got, got1, err := d.LoadClass(tt.args.className)
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

func TestDirClassLoader_ToString(t *testing.T) {
	type fields struct {
		absDir string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name:   "1",
			fields: fields{absDir: "C:\\Program Files\\Java\\jdk1.8.0_161\\jre\\lib"},
			want:   "DirClassLoader",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DirClassLoader{
				absDir: tt.fields.absDir,
			}
			if got := d.ToString(); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
