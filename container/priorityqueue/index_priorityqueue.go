package priorityqueue

import "container/heap"

func NewWithIndex[T any](cmp func(x, y T) bool) *IndexPriorityQuque[T] {
	return &IndexPriorityQuque[T]{impl: indexPriorityQueueImpl[T]{less: cmp}}
}

type IndexPriorityItem[T any] struct {
	index int
	Val   T
}

type IndexPriorityQuque[T any] struct {
	impl indexPriorityQueueImpl[T]
}

func (q IndexPriorityQuque[T]) Len() int {
	return q.impl.Len()
}

func (q *IndexPriorityQuque[T]) Push(val T) *IndexPriorityItem[T] {
	elem := &IndexPriorityItem[T]{
		index: -1,
		Val:   val,
	}

	heap.Push(&q.impl, elem)
	return elem
}

func (q *IndexPriorityQuque[T]) Pop() T {
	return heap.Pop(&q.impl).(*IndexPriorityItem[T]).Val
}

func (q *IndexPriorityQuque[T]) Remove(index int) T {
	return heap.Remove(&q.impl, index).(*IndexPriorityItem[T]).Val
}

func (q *IndexPriorityQuque[T]) Fix(index int) {
	heap.Fix(&q.impl, index)
}

func (q *IndexPriorityQuque[T]) FixItem(item *IndexPriorityItem[T]) {
	heap.Fix(&q.impl, item.index)
}

type indexPriorityQueueImpl[T any] struct {
	less  func(x, y T) bool
	items []*IndexPriorityItem[T]
}

func (q indexPriorityQueueImpl[T]) Len() int {
	return len(q.items)
}

func (q indexPriorityQueueImpl[T]) Swap(i, j int) {
	q.items[i], q.items[j] = q.items[j], q.items[i]
	q.items[i].index, q.items[j].index = i, j
}

func (q indexPriorityQueueImpl[T]) Less(i, j int) bool {
	return q.less(q.items[i].Val, q.items[j].Val)
}

func (q *indexPriorityQueueImpl[T]) Push(val any) {
	el := val.(*IndexPriorityItem[T])
	el.index = q.Len()
	q.items = append(q.items, el)
}

func (q *indexPriorityQueueImpl[T]) Pop() any {
	last := len(q.items) - 1
	el := q.items[last]
	q.items = q.items[:last]
	el.index = -1
	return el
}
