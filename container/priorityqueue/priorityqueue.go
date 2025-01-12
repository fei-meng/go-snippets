package priorityqueue

import "container/heap"

func New[T any](cmp func(x, y T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{impl: priorityQueueImpl[T]{less: cmp}}
}

type PriorityQueue[T any] struct {
	impl priorityQueueImpl[T]
}

func (q PriorityQueue[T]) Len() int {
	return q.impl.Len()
}

func (q *PriorityQueue[T]) Push(val T) {
	heap.Push(&q.impl, val)
}

func (q *PriorityQueue[T]) Pop() T {
	return heap.Pop(&q.impl).(T)
}

func (q *PriorityQueue[T]) Remove(index int) T {
	return heap.Remove(&q.impl, index).(T)
}

func (q *PriorityQueue[T]) Fix(index int) {
	heap.Fix(&q.impl, index)
}

type priorityQueueImpl[T any] struct {
	less  func(x, y T) bool
	items []T
}

func (q priorityQueueImpl[T]) Len() int {
	return len(q.items)
}

func (q priorityQueueImpl[T]) Swap(i, j int) {
	q.items[i], q.items[j] = q.items[j], q.items[i]
}

func (q priorityQueueImpl[T]) Less(i, j int) bool {
	return q.less(q.items[i], q.items[j])
}

func (q *priorityQueueImpl[T]) Push(val any) {
	q.items = append(q.items, val.(T))
}

func (q *priorityQueueImpl[T]) Pop() any {
	last := len(q.items) - 1
	val := q.items[last]
	q.items = q.items[:last]
	return val
}
