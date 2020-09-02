package classloader

import (
	"archive/zip"
	"reflect"
	"testing"
)

func TestCreateZipLoader(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *ZipClassLoader
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{path: "C:\\Program Files\\Java\\jdk1.8.0_161\\jre\\lib\\rt.jar"},
		},
		{
			name: "2",
			args: args{path: "C:\\Program Files\\Java\\jdk1.8.0_161\\jre\\lib\\aaa.jar"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateZipLoader(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateZipLoader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZipClassLoader_LoadClass(t *testing.T) {
	type fields struct {
		absZipPath string
		zipFile    *zip.ReadCloser
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
			path: "C:\\Program Files\\Java\\jdk1.8.0_161\\jre\\lib\\rt.jar",
			args: args{className: "java.lang.Class"},
		},
		{
			name: "2",
			path: "C:\\Program Files\\Java\\jdk1.8.0_161\\jre\\lib\\rt.jar",
			args: args{className: "java.lang.String"},
		},
		{
			name: "3",
			path: "C:\\Program Files\\Java\\jdk1.8.0_161\\jre\\lib\\rt.jar",
			args: args{className: "java.lang.zihua"},
		},
		{
			name: "4",
			path: "C:\\Program Files\\Java\\jdk1.8.0_161\\jre\\lib\\aaa.jar",
			args: args{className: "java.lang.String"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//z := ZipClassLoader{
			//	absZipPath: tt.fields.absZipPath,
			//}
			z := CreateZipLoader(tt.path)
			got, got1, err := z.LoadClass(tt.args.className)
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

func TestZipClassLoader_ToString(t *testing.T) {
	type fields struct {
		absZipPath string
		zipFile    *zip.ReadCloser
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name:   "1",
			fields: fields{},
			want:   "ZipClassLoader",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := ZipClassLoader{
				absZipPath: tt.fields.absZipPath,
			}
			if got := z.ToString(); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
