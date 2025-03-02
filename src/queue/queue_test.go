package queue_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "github.com/crissyro/data-structures/queue"
)

func TestQueue(t *testing.T) {
    t.Parallel()

    tests := []struct {
        name    string
        setup   func() *queue.Queue[string]
        ops     []func(*queue.Queue[string])
        wantErr bool
    }{
        {
            name: "Basic enqueue/dequeue",
            setup: func() *queue.Queue[string] {
                q := queue.New[string]()
                q.Enqueue("first")
                q.Enqueue("second")
                return q
            },
            ops: []func(*queue.Queue[string]){
                func(q *queue.Queue[string]) {
                    val, err := q.Dequeue()
                    require.NoError(t, err)
                    assert.Equal(t, "first", val)
                },
                func(q *queue.Queue[string]) {
                    val, err := q.Dequeue()
                    require.NoError(t, err)
                    assert.Equal(t, "second", val)
                },
            },
        },
        {
            name: "Empty queue handling",
            setup: func() *queue.Queue[string] {
                return queue.New[string]()
            },
            ops: []func(*queue.Queue[string]){
                func(q *queue.Queue[string]) {
                    _, err := q.Dequeue()
                    assert.Error(t, err, "Expected error on empty queue")
                },
            },
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()
            q := tt.setup()
            for _, op := range tt.ops {
                op(q)
            }
        })
    }
}

func TestQueueOrder(t *testing.T) {
    t.Parallel()

    q := queue.New[int]()
    expected := []int{1, 2, 3, 4, 5}

    for _, v := range expected {
        q.Enqueue(v)
    }

    for _, v := range expected {
        val, err := q.Dequeue()
        require.NoError(t, err)
        assert.Equal(t, v, val)
    }
}