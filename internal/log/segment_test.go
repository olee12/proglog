package log

import (
	"io"
	"io/ioutil"
	"os"
	"testing"

	api "github.com/olee12/proglog/api/v1"
	"github.com/stretchr/testify/require"
)

func TestSegment(t *testing.T) {
	dir, _ := ioutil.TempDir("", "segment-test")
	defer os.RemoveAll(dir)

	want := &api.Record{Value: []byte("hello world")}
	c := Config{}
	c.Segment.MaxStoreBytes = 1024
	c.Segment.MaxIndexByte = 3 * entWidth

	s, err := newSegment(dir, 16, c)
	require.NoError(t, err)
	require.Equal(t, uint64(16), s.nextOffset, s.baseOffset)
	require.False(t, s.IsMaxed())

	for i := uint64(0); i < 3; i++ {
		off, err := s.Append(want)
		require.NoError(t, err)
		require.Equal(t, 16+i, off)
		got, err := s.Read(off)
		require.NoError(t, err)
		require.Equal(t, want.Value, got.Value)
	}
	_, err = s.Append(want)
	require.Equal(t, io.EOF, err)
	// maxed index
	require.True(t, s.IsMaxed())

	c.Segment.MaxStoreBytes = uint64(3 * len(want.Value))
	c.Segment.MaxIndexByte = 1024

	s, err = newSegment(dir, 16, c)
	require.NoError(t, err)

	// maxed store
	require.True(t, s.IsMaxed())

	err = s.Close()
	require.NoError(t, err)

	s, err = newSegment(dir, 16, c)
	require.NoError(t, err)

	got, err := s.Read(16)
	require.NoError(t, err)
	require.Equal(t, want.Value, got.Value)

	err = s.Remove()
	require.NoError(t, err)

	s, err = newSegment(dir, 16, c)
	require.NoError(t, err)
	require.False(t, s.IsMaxed())
}
