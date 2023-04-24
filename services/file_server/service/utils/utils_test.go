package utils

import (
	"testing"
)

func TestRemoveSpecialChar(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{val: "abc_def(a).png"},
			want: "abcdefa.png",
		},
		{
			name: "Test 2",
			args: args{val: "Quick Brown fox Jumps.jpeg"},
			want: "QuickBrownfoxJumps.jpeg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveSpecialChar(tt.args.val); got != tt.want {
				t.Errorf("RemoveSpecialChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructPath(t *testing.T) {
	type args struct {
		basePath string
		blogId   string
		fileName string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			name:  "Test 1",
			args:  args{basePath: "/cloud", blogId: "2345", fileName: "abcd.jpg"},
			want:  "/cloud/2345",
			want1: "/cloud/2345/abcd.jpg",
		},
		{
			name:  "Test 2",
			args:  args{basePath: "/home/ubuntu/cloud", blogId: "abcdefdr", fileName: "xyx.png"},
			want:  "/home/ubuntu/cloud/abcdefdr",
			want1: "/home/ubuntu/cloud/abcdefdr/xyx.png",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ConstructPath(tt.args.basePath, tt.args.blogId, tt.args.fileName)
			if got != tt.want {
				t.Errorf("ConstructPath() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ConstructPath() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
