package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	s := Set[int]{}
	s.Add(1)
	require.True(t, s.Has(1), "set should contain 1 after adding")
}

func TestRemove(t *testing.T) {
	t.Parallel()
	s := Set[int]{}
	s.Add(2)
	s.Remove(2)
	require.False(t, s.Has(2), "set should not contain 2 after removal")
}

func TestHas(t *testing.T) {
	t.Parallel()
	s := Set[string]{}
	s.Add("hello")
	require.True(t, s.Has("hello"), "set should contain 'hello'")
	assert.False(t, s.Has("world"), "set should not contain 'world'")
}

func TestSetFrom(t *testing.T) {
	t.Parallel()
	s := SetFrom(1, 2, 3)
	for _, v := range []int{1, 2, 3} {
		assert.True(t, s.Has(v), "set should contain %d", v)
	}
	assert.False(t, s.Has(4), "set should not contain 4")
}

func BenchmarkAdd(b *testing.B) {
	s := make(Set[int])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
}

func BenchmarkRemove(b *testing.B) {
	s := make(Set[int])
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Remove(i)
	}
}

func BenchmarkHas(b *testing.B) {
	s := make(Set[int])
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.Has(i)
	}
}
