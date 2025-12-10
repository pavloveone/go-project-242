package code

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSize(t *testing.T) {
	tests := []struct {
		name  string
		path  string
		human bool
		all   bool
		want  string
	}{
		{
			name:  "regular file",
			path:  "./testdata/file1.txt",
			human: false,
			all:   false,
			want:  "9B\t./testdata/file1.txt",
		},
		{
			name:  "directory",
			path:  "./testdata/2",
			human: false,
			all:   false,
			want:  "48B\t./testdata/2",
		},
		{
			name:  "nested file",
			path:  "./testdata/2/file2.txt",
			human: false,
			all:   false,
			want:  "24B\t./testdata/2/file2.txt",
		},
		{
			name:  "human-readable",
			path:  "./testdata/bigFile.txt",
			human: true,
			all:   false,
			want:  "9.5KB\t./testdata/bigFile.txt",
		},
		{
			name:  "hidden file excluded",
			path:  "./testdata/.secret",
			human: true,
			all:   false,
			want:  "0B\t./testdata/.secret",
		},
		{
			name:  "hidden file included",
			path:  "./testdata/.secret",
			human: true,
			all:   true,
			want:  "200.0KB\t./testdata/.secret",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			got, err := GetSize(tt.path, tt.human, tt.all)
			r.NoError(err)
			r.Equal(tt.want, got)
		})
	}
}

func TestFormatSize(t *testing.T) {
	tests := []struct {
		name  string
		size  int64
		human bool
		want  string
		err   bool
	}{
		{
			name:  "human bytes",
			size:  123,
			human: true,
			want:  "123B",
		}, {
			name:  "human MB",
			size:  25165824,
			human: true,
			want:  "24.0MB",
		},
		{
			name:  "negative size",
			size:  -1,
			human: true,
			err:   true,
		},
		{
			name:  "raw bytes",
			size:  25165824,
			human: false,
			want:  "25165824B",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)
			got, err := FormatSize(tt.size, tt.human)
			if tt.err {
				r.Error(err)
				return
			}
			r.NoError(err)
			r.Equal(tt.want, got)
		})
	}
}
