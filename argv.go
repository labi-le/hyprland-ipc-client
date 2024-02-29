// thanks https://www.quasilyte.dev/blog/post/pathfinding/#the-priority-queue

package client

import "math/bits"

type PriorityQueue[T any] struct {
	buckets [128][]T
	mask    uint64

	forward int
	back    int
	own     int
}

func NewQueue[T any]() *PriorityQueue[T] {
	own := 64 / 2
	return &PriorityQueue[T]{
		forward: own,
		back:    own - 1,
		own:     own,
	}
}

func (q *PriorityQueue[T]) Push(priority int, value T) {
	// A q.buckets[i] boundcheck is removed due to this &-masking.
	i := uint(priority) & 0b111111
	q.buckets[i] = append(q.buckets[i], value)
	q.mask |= 1 << i
}

func (q *PriorityQueue[T]) Pop() T {
	// The TrailingZeros64 on amd64 is a couple of
	// machine instructions (BSF and CMOV).
	//
	// We only need to execute these two to
	// get an index of a non-empty bucket.
	i := uint(bits.TrailingZeros64(q.mask))

	// This explicit length check is needed to remove
	// the q.buckets[i] bound checks below.
	if i < uint(len(q.buckets)) {
		// These two lines perform a pop operation.
		e := q.buckets[i][len(q.buckets[i])-1]
		q.buckets[i] = q.buckets[i][:len(q.buckets[i])-1]

		if len(q.buckets[i]) == 0 {
			// If this bucket is empty now, clear the associated bit.
			// This preserves the invariant.
			q.mask &^= 1 << i
		}
		return e
	}

	// If the queue is empty, return a zero value.
	var x T
	return x
}

func (q *PriorityQueue[T]) Reset() {
	mask := q.mask

	// The first bucket to clear is the first non-empty one.
	// Skip all empty buckets "to the right".
	offset := uint(bits.TrailingZeros64(mask))
	mask >>= offset
	i := offset

	// When every bucket "to the left" is empty, the mask will be
	// equal to 0 and this loop will terminate.
	for mask != 0 {
		q.buckets[i] = q.buckets[i][:0]
		mask >>= 1
		i++
	}

	q.mask = 0
}

func (q *PriorityQueue[T]) IsEmpty() bool {
	return q.mask == 0
}

func (q *PriorityQueue[T]) Len() int {
	return bits.OnesCount64(q.mask)
}

func (q *PriorityQueue[T]) Peek() T {
	i := uint(bits.TrailingZeros64(q.mask))
	return q.buckets[i][len(q.buckets[i])-1]
}

func (q *PriorityQueue[T]) Add(value T) {
	if q.forward == q.Len() {
		panic("queue is full")
	}
	q.Push(q.forward, value)
	q.forward++
}

func (q *PriorityQueue[T]) Back(value T) {
	if q.back == q.own {
		panic("queue is full")
	}
	q.Push(q.back, value)
	q.back--
}

type ByteQueue struct {
	*PriorityQueue[[]byte]
}

func NewByteQueue() *ByteQueue {
	return &ByteQueue{PriorityQueue: NewQueue[[]byte]()}
}

func (q *ByteQueue) Glue() []byte {
	var glued []byte

	for !q.IsEmpty() {
		glued = append(glued, q.Pop()...)
		glued = append(glued, ' ')
	}

	if len(glued) > 0 {
		return glued[:len(glued)-1]
	}
	return glued
}
