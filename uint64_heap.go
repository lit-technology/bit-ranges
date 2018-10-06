package bits

// ItemUint64 is a Key Value for uint64.
type ItemUint64 struct {
	key   uint64
	value uint64
	index int // The index of the item in the heap
}

// NewItemUint64 create a new ItemUint64 with Key Value
func NewItemUint64(key, value uint64) *ItemUint64 {
	return &ItemUint64{
		key:   key,
		value: value,
	}
}

// HeapUint64 is a Heap for uint64.
type HeapUint64 []*ItemUint64

// Length of the Heap
func (h HeapUint64) Len() int { return len(h) }

// Less
func (h HeapUint64) Less(i, j int) bool {
	return h[i].key > h[j].key
}

// Swap Index i, j in the Heap.
func (h HeapUint64) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

// Push an ItemUint64 and Heapify.
func (h *HeapUint64) Push(item *ItemUint64) {
	n := len(*h)
	item.index = n
	*h = append(*h, item)
	h.up(h.Len() - 1)
}

// Pop nth Heap Item and Heapify 0..n-1.
func (h *HeapUint64) Pop() *ItemUint64 {
	n := h.Len() - 1
	h.Swap(0, n)
	h.down(0, n)

	old := *h
	n = len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*h = old[0 : n-1]
	return item
}

// Heapify Element Up
func (h *HeapUint64) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

// Heapify Element Down
func (h *HeapUint64) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}
