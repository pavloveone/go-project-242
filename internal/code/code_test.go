package code

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSize(t *testing.T) {
	r := require.New(t)
	first := "./testdata/file1.txt"
	second := "./testdata/2"
	third := "./testdata/2/file2.txt"
	fourth := "./testdata/bigFile.txt"

	res1, err := GetSize(first, false)
	r.NoError(err)
	r.Equal(fmt.Sprintf("9B\t%s", first), res1)

	res2, err := GetSize(second, false)
	r.NoError(err)
	r.Equal(fmt.Sprintf("48B\t%s", second), res2)

	res3, err := GetSize(third, false)
	r.NoError(err)
	r.Equal(fmt.Sprintf("24B\t%s", third), res3)

	res4, err := GetSize(fourth, true)
	r.NoError(err)
	r.Equal(fmt.Sprintf("9.5KB\t%s", fourth), res4)
}

func TestFormatSize(t *testing.T) {
	r := require.New(t)
	res1, err := FormatSize(1234567, true)
	r.NoError(err)
	r.Equal(res1, "1.2MB")

	res2, err := FormatSize(123, true)
	r.NoError(err)
	r.Equal(res2, "123B")

	res3, _ := FormatSize(-1, true)
	r.NoError(err)
	r.Equal(res3, "")

	res4, err := FormatSize(25165824, false)
	r.NoError(err)
	r.Equal(res4, "25165824B")
}
