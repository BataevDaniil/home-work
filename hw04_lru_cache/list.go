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
	front *ListItem
	back  *ListItem
	len   int
}

func (l *list) Len() (count int) {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(item interface{}) *ListItem {
	listItem := &ListItem{Value: item, Next: l.front}
	if l.front == nil {
		l.back = listItem
	} else {
		l.front.Prev = listItem
	}
	l.front = listItem
	l.len++
	return listItem
}

func (l *list) PushBack(item interface{}) *ListItem {
	listItem := &ListItem{Value: item, Prev: l.back}
	if l.back == nil {
		l.front = listItem
	} else {
		l.back.Next = listItem
	}
	l.back = listItem
	l.len++
	return listItem
}

func (l *list) Remove(item *ListItem) {
	if l.front == item {
		l.front = item.Next
	} else {
		item.Prev.Next = item.Next
	}
	if l.back == item {
		l.back = item.Prev
	} else {
		item.Next.Prev = item.Prev
	}
	l.len--
}

func (l *list) MoveToFront(item *ListItem) {
	l.Remove(item)
	l.PushFront(item.Value)
}

func NewList() List {
	return new(list)
}
