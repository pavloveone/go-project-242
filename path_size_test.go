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
			human:     false,
			all:       false,
			recursive: false,
			want:      "9B",
		},
		{
			name:      "directory",
			path:      "./testdata/2",
			human:     false,
			all:       false,
			recursive: false,
			want:      "48B",
		},
		{
			name:      "nested file",
			path:      "./testdata/2/file2.txt",
			human:     false,
			all:       false,
			recursive: false,
			want:      "24B",
		},
		{
			name:      "human-readable",
			path:      "./testdata/bigFile.txt",
			human:     true,
			all:       false,
			recursive: false,
			want:      "9.5KB",
		},
		{
			name:      "hidden file excluded",
			path:      "./testdata/.secret",
			human:     true,
			all:       false,
			recursive: false,
			want:      "0B",
		},
		{
			name:      "hidden file included",
			path:      "./testdata/.secret",
			human:     true,
			all:       true,
			recursive: false,
			want:      "200.0KB",
		},
		{
			name:      "recursion with exclusion of hidden files and directories",
			path:      "./testdata/",
			human:     true,
			all:       false,
			recursive: true,
			want:      "344.9KB",
		},
		{
			name:      "recursion with inclusion of hidden files and directories",
			path:      "./testdata/",
			human:     true,
			all:       true,
			recursive: true,
			want:      "611.5KB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			got, err := GetPathSize(tt.path, tt.human, tt.all, tt.recursive)
			r.NoError(err)
			r.Equal(tt.want, got)
		})
	}
}
