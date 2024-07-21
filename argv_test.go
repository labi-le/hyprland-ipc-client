package client

import (
	"bytes"
	"testing"
)

func Test_priorityQueue_Add(t *testing.T) {
	queue := NewByteQueue()
	queue.batch = true
	queue.command = []byte("command")
	queue.Add([]byte("name"))
	queue.Add([]byte("ivan"))
	queue.Back([]byte("my"))
	queue.Back([]byte("hello"))

	glued := queue.Glue()
	if !bytes.Equal(glued, []byte("[[BATCH]]command hello;command my;command name;command ivan")) {
		t.Errorf("got %s", glued)
	}
}
