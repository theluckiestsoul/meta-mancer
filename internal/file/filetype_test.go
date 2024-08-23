package file

import "testing"

func Test_getFileType(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with a JPEG file",
			args: args{
				filePath: "IMG_4229.JPG",
			},
			want: "image/jpeg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFileType(tt.args.filePath); got != tt.want {
				t.Errorf("getFileType() = %v, want %v", got, tt.want)
			}
		})
	}
}
