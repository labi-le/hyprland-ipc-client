package client

import (
	"bytes"
	"testing"
)

func Test_priorityQueue_Add(t *testing.T) {
	queue := NewByteQueue()
	queue.Add([]byte("name"))
	queue.Add([]byte("ivan"))
	queue.Back([]byte("my"))
	queue.Back([]byte("hello"))

	glued := queue.Glue()
	if !bytes.Equal(glued, []byte("hello my name ivan")) {
		t.Errorf("got %s", glued)
	}
}
