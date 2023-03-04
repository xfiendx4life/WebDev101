package main

import (
	"context"
	"testing"
)

func TestQueuePush(t *testing.T) {
	q := newQueue("test_queue")
	err := q.push(context.Background(), "test_message")
	if err != nil {
		t.Fatal(err)
	} else if q.data[0] != "test_message" {
		t.Fail()
	}
}

func TestQueuePop(t *testing.T) {
	q := newQueue("test_queue")
	q.data = append(q.data, "test_message")
	res, err := q.pop(context.Background())
	if err != nil {
		t.Fatal(err)
	} else if res != "test_message" {
		t.Fail()
	}
}

func TestQueuePopError(t *testing.T) {
	q := newQueue("test_queue")
	// q.data = append(q.data, "test_message")
	res, err := q.pop(context.Background())
	if err == nil {
		t.Fatal(err)
	} else if res != "" {
		t.Fail()
	}
}
