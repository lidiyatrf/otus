package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}

func TestLen(t *testing.T) {
	l := NewList()

	require.Equal(t, 0, l.Len())
	l.PushFront(3) // [3]
	require.Equal(t, 1, l.Len())
	l.PushBack(4) // [3 4]
	require.Equal(t, 2, l.Len())
	l.PushFront(2) // [2 3 4]
	require.Equal(t, 3, l.Len())
	l.PushBack(5) // [2 3 4 5]
	require.Equal(t, 4, l.Len())

	l.Remove(l.Front())
	require.Equal(t, 3, l.Len())
	l.Remove(l.Back())
	require.Equal(t, 2, l.Len())
	l.Remove(l.Front())
	require.Equal(t, 1, l.Len())
	l.Remove(l.Back())
	require.Equal(t, 0, l.Len())
}

func TestPushFront(t *testing.T) {
	l := NewList()

	require.Nil(t, l.Front())

	item := 3
	l.PushFront(item) // [3]
	require.Equal(t, item, l.Front().Value)

	item = 5
	l.PushFront(item) // [5 3]
	require.Equal(t, item, l.Front().Value)

	item = 7
	l.PushFront(item) // [7 5 3]
	require.Equal(t, item, l.Front().Value)
}

func TestPushBack(t *testing.T) {
	l := NewList()

	require.Nil(t, l.Back())

	item := 3
	l.PushBack(item) // [3]
	require.Equal(t, item, l.Back().Value)

	item = 5
	l.PushBack(item) // [3 5]
	require.Equal(t, item, l.Back().Value)

	item = 7
	l.PushBack(item) // [3 5 7]
	require.Equal(t, item, l.Back().Value)
}

func TestMoveToFront(t *testing.T) {
	t.Run("move to front head", func(t *testing.T) {
		l := NewList()

		l.PushFront(1)
		l.PushFront(2)
		l.PushFront(3) // [3 2 1]
		require.Equal(t, 3, l.Len())

		l.MoveToFront(l.Front()) // [3 2 1]
		require.Equal(t, 3, l.Len())

		require.Equal(t, 3, l.Front().Value)
		l.Remove(l.Front())
		require.Equal(t, 2, l.Front().Value)
		l.Remove(l.Front())
		require.Equal(t, 1, l.Front().Value)
		l.Remove(l.Front())
		require.Nil(t, l.Front())
	})

	t.Run("move to front tail", func(t *testing.T) {
		l := NewList()

		l.PushFront(1)
		l.PushFront(2)
		l.PushFront(3) // [3 2 1]
		require.Equal(t, 3, l.Len())

		l.MoveToFront(l.Back()) // [1 3 2]
		require.Equal(t, 3, l.Len())

		require.Equal(t, 1, l.Front().Value)
		l.Remove(l.Front())
		require.Equal(t, 3, l.Front().Value)
		l.Remove(l.Front())
		require.Equal(t, 2, l.Front().Value)
		l.Remove(l.Front())
		require.Nil(t, l.Front())
	})

	t.Run("move to front middle", func(t *testing.T) {
		l := NewList()

		l.PushFront(1)
		l.PushFront(2)
		l.PushFront(3) // [3 2 1]
		require.Equal(t, 3, l.Len())

		l.MoveToFront(l.Front().Next) // [2 3 1]
		require.Equal(t, 3, l.Len())

		require.Equal(t, 2, l.Front().Value)
		l.Remove(l.Front())
		require.Equal(t, 3, l.Front().Value)
		l.Remove(l.Front())
		require.Equal(t, 1, l.Front().Value)
		l.Remove(l.Front())
		require.Nil(t, l.Front())
	})

	t.Run("move to front front in the list of size = 1", func(t *testing.T) {
		l := NewList()

		l.PushBack(1) // [1]
		require.Equal(t, 1, l.Len())

		l.MoveToFront(l.Front()) // [1]
		require.Equal(t, 1, l.Len())

		require.Equal(t, 1, l.Front().Value)
		l.Remove(l.Front())
		require.Nil(t, l.Front())
	})

	t.Run("move to front back in the list of size = 1", func(t *testing.T) {
		l := NewList()

		l.PushBack(1) // [1]
		require.Equal(t, 1, l.Len())

		l.MoveToFront(l.Back()) // [1]
		require.Equal(t, 1, l.Len())

		require.Equal(t, 1, l.Front().Value)
		l.Remove(l.Front())
		require.Nil(t, l.Front())
	})
}

func TestRemove(t *testing.T) {
	t.Run("removing from the head", func(t *testing.T) {
		l := NewList()

		l.PushFront(1)
		l.PushFront(2)
		l.PushFront(3) // [3 2 1]

		l.Remove(l.Front()) // [2 1]

		require.Equal(t, 2, l.Front().Value)
		l.Remove(l.Front())
		require.Equal(t, 1, l.Front().Value)
		l.Remove(l.Front())
		require.Nil(t, l.Front())
	})

	t.Run("removing from the tail", func(t *testing.T) {
		l := NewList()

		l.PushFront(1)
		l.PushFront(2)
		l.PushFront(3) // [3 2 1]

		l.Remove(l.Back()) // [3 2]

		require.Equal(t, 3, l.Front().Value)
		l.Remove(l.Front())
		require.Equal(t, 2, l.Front().Value)
		l.Remove(l.Front())
		require.Nil(t, l.Front())
	})

	t.Run("removing from the middle", func(t *testing.T) {
		l := NewList()

		l.PushFront(1)
		l.PushFront(2)
		l.PushFront(3) // [3 2 1]

		l.Remove(l.Back().Prev) // [3 1]

		require.Equal(t, 3, l.Front().Value)
		l.Remove(l.Front())
		require.Equal(t, 1, l.Front().Value)
		l.Remove(l.Front())
		require.Nil(t, l.Front())
	})

	t.Run("removing back from the list of size = 1", func(t *testing.T) {
		l := NewList()

		l.PushFront(1) // [1]
		require.Equal(t, 1, l.Len())
		l.Remove(l.Back()) // []
		require.Equal(t, 0, l.Len())

		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("removing front from the list of size = 1", func(t *testing.T) {
		l := NewList()

		l.PushFront(1) // [1]
		require.Equal(t, 1, l.Len())
		l.Remove(l.Front()) // []
		require.Equal(t, 0, l.Len())

		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})
}
