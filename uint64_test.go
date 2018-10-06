package bits

import (
	"math/bits"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeUint64(t *testing.T) {
	for i := uint64(1); i < 2; i++ {
		for bit := i; bit > 0; bit-- {
			for n := uint64(0); n < 1<<4; n++ {
				start, end := RangeUint64(bit, n<<bit)
				// TODO: Test with less false positives.
				assert.True(t, bits.TrailingZeros(uint(start)) >= int(bit))
				assert.Equal(t, 0, bits.TrailingZeros(uint(end)))
			}
		}
	}
}

func BenchmarkRangeUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for n := uint64(512); i < 1024; i += 22 {
			for bits := uint64(1); bits < 6; bits++ {
				_, _ = RangeUint64(bits, n)
			}
		}
	}
}

func TestRangesUint64(t *testing.T) {
	n1 := uint64(1 << 48)
	n2 := uint64(1 << 56)
	ns := make([]uint64, 0)
	for n := n1; n <= n2; n += (1 << 48) {
		ns = append(ns, n)
	}
	assert.Equal(t, 2, len(RangesUint64(48, ns...)))
	ns = make([]uint64, 0)
	for n := n1; n <= n2; n += (1 << 40) {
		ns = append(ns, n)
	}
	assert.Equal(t, 2, len(RangesUint64(40, ns...)))
}

func TestMergeHeap(t *testing.T) {
	h := new(HeapUint64)
	h.Push(NewItemUint64(0, 99))
	h.Push(NewItemUint64(100, 199))
	h.Push(NewItemUint64(300, 399))
	assert.Equal(t, 4, len(PopToArrayRanges(h)))
	assert.Equal(t, 0, len(PopToArrayRanges(h)))

	h.Push(NewItemUint64(100, 199))
	h.Push(NewItemUint64(0, 99))
	h.Push(NewItemUint64(300, 399))
	h.Push(NewItemUint64(200, 350))
	assert.Equal(t, 2, len(PopToArrayRanges(h)))
	assert.Equal(t, 0, len(PopToArrayRanges(h)))

	h.Push(NewItemUint64(100, 350))
	h.Push(NewItemUint64(300, 399))
	h.Push(NewItemUint64(100, 350))
	assert.Equal(t, 2, len(PopToArrayRanges(h)))
	assert.Equal(t, 0, len(PopToArrayRanges(h)))
}
