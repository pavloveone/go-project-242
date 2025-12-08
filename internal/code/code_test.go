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

	res1, err := GetSize(first)
	r.NoError(err)
	r.Equal(fmt.Sprintf("9\t%s", first), res1)

	res2, err := GetSize(second)
	r.NoError(err)
	r.Equal(fmt.Sprintf("48\t%s", second), res2)

	res3, err := GetSize(third)
	r.NoError(err)
	r.Equal(fmt.Sprintf("24\t%s", third), res3)
}
