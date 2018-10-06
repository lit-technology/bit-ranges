package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUint64Heap(t *testing.T) {
	h := make(HeapUint64, 0)
	h.Push(NewItemUint64(2, 2))
	h.Push(NewItemUint64(1, 1))
	h.Push(NewItemUint64(3, 3))
	assert.Equal(t, h.Len(), 3)
	for i := uint64(3); i > 0; i-- {
		item := h.Pop()
		assert.Equal(t, i, item.key, "Key is incorrect")
		assert.Equal(t, i, item.value, "Value is incorrect")
	}
	assert.Empty(t, h)
}
