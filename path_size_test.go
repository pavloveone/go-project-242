package code

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSize(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		human     bool
		all       bool
		recursive bool
		want      string
	}{
		{
			name:      "regular file",
			path:      "./testdata/file1.txt",
			recursive: false,
			human:     false,
			all:       false,
			want:      "9B",
		},
		{
			name:      "directory",
			path:      "./testdata/2",
			recursive: false,
			human:     false,
			all:       false,
			want:      "48B",
		},
		{
			name:      "nested file",
			path:      "./testdata/2/file2.txt",
			recursive: false,
			human:     false,
			all:       false,
			want:      "24B",
		},
		{
			name:      "human-readable",
			path:      "./testdata/bigFile.txt",
			recursive: false,
			human:     true,
			all:       false,
			want:      "9.5KB",
		},
		{
			name:      "hidden file excluded",
			path:      "./testdata/.secret",
			recursive: false,
			human:     true,
			all:       false,
			want:      "0B",
		},
		{
			name:      "hidden file included",
			path:      "./testdata/.secret",
			recursive: false,
			human:     true,
			all:       true,
			want:      "200.0KB",
		},
		{
			name:      "recursion with exclusion of hidden files and directories",
			path:      "./testdata/",
			recursive: true,
			human:     true,
			all:       false,
			want:      "344.9KB",
		},
		{
			name:      "recursion with inclusion of hidden files and directories",
			path:      "./testdata/",
			recursive: true,
			human:     true,
			all:       true,
			want:      "611.5KB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			got, err := GetPathSize(tt.path, tt.recursive, tt.human, tt.all)
			r.NoError(err)
			r.Equal(tt.want, got)
		})
	}
}
