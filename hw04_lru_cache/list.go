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
	len   int
	front *ListItem
	back  *ListItem
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	oldFront := l.Front()
	newFront := &ListItem{v, oldFront, nil}

	if oldFront != nil {
		oldFront.Prev = newFront
	}

	l.front = newFront

	if l.Back() == nil {
		l.back = newFront
	}

	l.len++

	return newFront
}

func (l *list) PushBack(v interface{}) *ListItem {
	oldBack := l.Back()
	newBack := &ListItem{v, nil, oldBack}

	if oldBack != nil {
		oldBack.Next = newBack
	}

	l.back = newBack

	if l.Front() == nil {
		l.front = newBack
	}

	l.len++

	return newBack
}

func (l *list) Remove(i *ListItem) {
	prev := i.Prev
	next := i.Next

	if prev != nil {
		prev.Next = next
	}

	if next != nil {
		next.Prev = prev
	}

	if l.Front() == i && next != nil {
		l.front = next
	}

	if l.Back() == i && prev != nil {
		l.back = prev
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}
