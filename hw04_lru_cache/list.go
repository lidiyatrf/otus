package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	head *ListItem
	tail *ListItem
	size int
}

func NewList() List {
	return new(list)
}

func (l list) Len() int {
	return l.size
}

func (l list) Front() *ListItem {
	return l.head
}

func (l list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	if l.head == nil {
		l.head = &ListItem{Prev: nil, Next: nil, Value: v}
		l.tail = l.head
		l.size = 1
		return l.head
	}

	l.head = &ListItem{Prev: nil, Next: l.head, Value: v}
	l.head.Next.Prev = l.head
	l.size++
	return l.head
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l.tail == nil {
		l.tail = &ListItem{Prev: nil, Next: nil, Value: v}
		l.head = l.tail
		l.size = 1
		return l.tail
	}

	l.tail = &ListItem{Prev: l.tail, Next: nil, Value: v}
	l.tail.Prev.Next = l.tail
	l.size++
	return l.tail
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i == l.head {
		l.head = l.head.Next
	}
	if i == l.tail {
		l.tail = l.tail.Prev
	}
	l.size--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)

	if l.head == nil {
		l.head = i
		l.tail = l.head
		l.size = 1
		return
	}

	i.Next = l.head
	i.Prev = nil
	l.head = i
	l.head.Next.Prev = l.head
	l.size++
}
