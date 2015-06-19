package heap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	h := new(IntHeap)
	h.Direction = Ascending
	h.Insert(30)
	h.Insert(10)
	h.Insert(50)

	_, err := h.Max()
	assert.NotNil(t, err, "Max() operation is not supported")

	min, err := h.Min()

	assert.Nil(t, err, "Extraction of h.Min()")
	assert.Equal(t, 10, min, "Min value from Heap")

	h.Insert(5)
	min, err = h.Min()
	assert.Nil(t, err, "Extraction of h.Min()")
	assert.Equal(t, 5, min, "Min value from Heap")

}

func TestParentIndex(t *testing.T) {
	h := new(IntHeap)

	equal := func(idx, result int) {
		assert.Equal(
			t,
			result,
			h.parentIndex(idx),
			fmt.Sprintf("Parent of %d", idx),
		)
	}

	equal(0, 0)
	equal(1, 0)
	equal(2, 0)
	equal(3, 1)
	equal(4, 1)
	equal(5, 2)
	equal(6, 2)
	equal(7, 3)
	equal(8, 3)
	equal(9, 4)
	equal(10, 4)
	equal(11, 5)
	equal(12, 5)
	equal(13, 6)
	equal(14, 6)
	equal(18, 8)

}

func TestChildIndexes(t *testing.T) {
	h := new(IntHeap)

	equal := func(idx, left, right int) {
		cleft, cright := h.childIndexes(idx)
		assert.Equal(t, left, cleft, fmt.Sprintf("Left Child of %d", idx))
		assert.Equal(t, right, cright, fmt.Sprintf("Right Child of %d", idx))
	}

	equal(0, 1, 2)
	equal(1, 3, 4)
	equal(2, 5, 6)
	equal(8, 17, 18)
}

func TestExtractMin(t *testing.T) {
	h := new(IntHeap)
	h.Direction = Ascending
	h.Insert(30)
	h.Insert(10)
	h.Insert(50)
	h.Insert(18)
	h.Insert(1)

	test := func(v int) {
		val, err := h.Extract()
		assert.Nil(t, err, "Extract() has failed")
		assert.Equal(t, val, v, "Extract() gives wrong result (min)")
	}

	test(1)
	test(10)
	test(18)
	test(30)
	test(50)
}

func TestExtractMax(t *testing.T) {
	h := new(IntHeap)
	h.Direction = Descending
	h.Insert(30)
	h.Insert(10)
	h.Insert(50)
	h.Insert(18)
	h.Insert(1)

	test := func(v int) {
		val, err := h.Extract()
		assert.Nil(t, err, "Extract() has failed")
		assert.Equal(t, val, v, "Extract() gives wrong result (max)")
	}

	test(50)
	test(30)
	test(18)
	test(10)
	test(1)

}
