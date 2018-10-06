package bits

// RangeUint64 creates a uint64 range from unsetting and setting n least significant bits.
func RangeUint64(bits uint64, n uint64) (uint64, uint64) {
	mask := uint64((1 << (bits)) - 1)
	start := n &^ mask
	end := n | mask
	return start, end
}

// RangesUint64 creates multiple uint64 ranges from unsetting and setting n least significant bits.
func RangesUint64(bits uint64, ns ...uint64) []uint64 {
	h := new(HeapUint64)
	for _, n := range ns {
		start, end := RangeUint64(bits, n)
		h.Push(NewItemUint64(start, end))
	}
	return PopToArrayRanges(h)
}

// PopToArrayRanges pops Heap Key Values as Ranges into Array and merge connecting ranges together.
func PopToArrayRanges(h *HeapUint64) []uint64 {
	if h.Len() == 0 {
		return make([]uint64, 0)
	}
	items := make([]uint64, h.Len()*2)
	count := 0 // For last 2 items.
	var prevItem = h.Pop()
	for h.Len() > 0 {
		item := h.Pop()
		if prevItem.key-1 <= item.value {
			prevItem.key = item.key
		} else if !(item.key >= prevItem.key && item.value <= prevItem.value) {
			items[count] = prevItem.value
			items[count+1] = prevItem.key
			count += 2
			prevItem = item
		}
	}
	items[count] = prevItem.value
	items[count+1] = prevItem.key
	count += 2
	return items[0:count]
}
