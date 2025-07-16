package doublylinkedlist

type DLNode struct {
	Value int
	Prev  *DLNode
	Next  *DLNode
}

type List struct {
	head *DLNode
	tail *DLNode
	size int
}

func New() *List {
	return &List{}
}

func (l *List) InsertFront(value int) {
	node := &DLNode{Value: value}

	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.Next = l.head
		l.head.Prev = node
		l.head = node
	}

	l.size++
}

func (l *List) InsertBack(value int) {
	node := &DLNode{Value: value}

	if l.tail == nil {
		l.head = node
		l.tail = node
	} else {
		node.Prev = l.tail
		l.tail.Next = node
		l.tail = node
	}

	l.size++
}

func (l *List) Delete(value int) bool {
	curr := l.head

	for curr != nil {
		if curr.Value == value {
			if curr.Prev != nil {
				curr.Prev.Next = curr.Next
			} else {
				l.head = curr.Next
			}

			if curr.Next != nil {
				curr.Next.Prev = curr.Prev
			} else {
				l.tail = curr.Prev
			}

			l.size--
			return true
		}
		curr = curr.Next
	}

	return false
}

func (l *List) Search(value int) *DLNode {
	for curr := l.head; curr != nil; curr = curr.Next {
		if curr.Value == value {
			return curr
		}
	}
	return nil
}

func (l *List) ToSlice() []int {
	var result []int
	for curr := l.head; curr != nil; curr = curr.Next {
		result = append(result, curr.Value)
	}
	return result
}

func (l *List) ToReverseSlice() []int {
	var result []int
	for curr := l.tail; curr != nil; curr = curr.Prev {
		result = append(result, curr.Value)
	}
	return result
}

func (l *List) Len() int {
	return l.size
}

func (l *List) IsEmpty() bool {
	return l.size == 0
}
