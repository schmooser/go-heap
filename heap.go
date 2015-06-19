package heap

import (
	"errors"
	//"log"
)

type Direction int

const (
	Ascending Direction = iota
	Descending
)

type Heap interface {
	Insert() error
	ExtractMin() error
	ExtractMax() error
}

type IntHeap struct {
	// Direction defines whether heap is ascending or descending
	Direction Direction

	// values holds actual values of the tree
	values []int
}

func (h *IntHeap) Min() (int, error) {
	if h.Direction != Ascending {
		return 0, errors.New("Heap should be of Ascending type to support Min operation")
	}
	if len(h.values) == 0 {
		return 0, errors.New("No min element present")
	}
	return h.values[0], nil
}

func (h *IntHeap) Max() (int, error) {
	if h.Direction != Descending {
		return 0, errors.New("Heap should be of Descending type to support Max operation")
	}
	if len(h.values) == 0 {
		return 0, errors.New("No max element present")
	}
	return h.values[0], nil
}

func (h *IntHeap) Extract() (int, error) {

	if h.values == nil {
		return 0, errors.New("No more elements in heap")
	}
	var val int
	var err error
	if h.Direction == Ascending {
		val, err = h.Min()
	} else {
		val, err = h.Max()
	}

	if err != nil {
		return 0, err
	}
	h.swap(0, len(h.values)-1)
	h.values = h.values[:len(h.values)-1]

	var idx, child, left, right int // indexes

	for {
		left, right = h.childIndexes(idx)
		if left >= len(h.values) && right >= len(h.values) {
			return val, nil
		}

		switch {
		case left < len(h.values) && right == len(h.values):
			child = left
		case h.cmp(h.values[right], h.values[left]):
			child = left
		default:
			child = right
		}

		if h.invariant(child) {
			return val, nil
		} else {
			h.swap(idx, child)
			idx = child
		}
	}

	return val, nil
}

// parentIndex returns index of the parent item for provided item index i.
func (h *IntHeap) parentIndex(i int) int {
	if i == 0 {
		return 0
	} else {
		return (i+1)>>1 - 1
	}
}

func (h *IntHeap) childIndexes(i int) (left, right int) {
	left = i<<1 + 1
	right = left + 1
	return
}

// Insert inserts provided x in the heap h.
func (h *IntHeap) Insert(x int) error {
	h.values = append(h.values, x)
	//defer log.Printf("array = %v", h.values)

	idx := len(h.values) - 1
	for !h.invariant(idx) {
		//log.Printf("iterating - array = %v", h.values)
		h.swap(idx, h.parentIndex(idx))
		idx = h.parentIndex(idx)
	}
	return nil
}

// swap swaps i'th and j'th elements
func (h *IntHeap) swap(i, j int) error {
	h.values[i], h.values[j] = h.values[j], h.values[i]
	return nil
}

// invariant checks whether heap invariant is satisfied for index i and it's
// parent.
func (h *IntHeap) invariant(i int) bool {
	return h.cmp(h.values[i], h.values[h.parentIndex(i)])
}

// cmp compares x and y
func (h *IntHeap) cmp(x, y int) bool {
	if h.Direction == Ascending {
		return x >= y
	} else {
		return y >= x
	}
}
