package stack_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "github.com/crissyro/data-structures/stack"
)

func TestStack(t *testing.T) {
    t.Parallel()

    tests := []struct {
        name     string
        initFunc func() *stack.Stack[int]
        ops      []func(*stack.Stack[int])
        wantErr  bool
    }{
        {
            name: "Empty stack",
            initFunc: func() *stack.Stack[int] {
                return stack.New[int]()
            },
            ops: []func(*stack.Stack[int]){
                func(s *stack.Stack[int]) {
                    assert.True(t, s.IsEmpty(), "New stack should be empty")
                    assert.Equal(t, uint64(0), s.Size(), "Size should be 0")
                },
            },
        },
        {
            name: "Push/Pop sequence",
            initFunc: func() *stack.Stack[int] {
                s := stack.New[int]()
                s.Push(1)
                s.Push(2)
                return s
            },
            ops: []func(*stack.Stack[int]){
                func(s *stack.Stack[int]) {
                    val, err := s.Pop()
                    require.NoError(t, err)
                    assert.Equal(t, 2, val)
                },
                func(s *stack.Stack[int]) {
                    val, err := s.Pop()
                    require.NoError(t, err)
                    assert.Equal(t, 1, val)
                },
                func(s *stack.Stack[int]) {
                    _, err := s.Pop()
                    assert.Error(t, err, "Expected error on empty stack")
                },
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()
            s := tt.initFunc()
            for _, op := range tt.ops {
                op(s)
            }
        })
    }
}

func TestStackConcurrency(t *testing.T) {
    t.Parallel()

    s := stack.New[int]()
    const numOps = 1000

    t.Run("Parallel pushes", func(t *testing.T) {
        t.Parallel()
        for i := 0; i < numOps; i++ {
            s.Push(i)
        }
    })

    t.Run("Parallel pops", func(t *testing.T) {
        t.Parallel()
        for i := 0; i < numOps; i++ {
            _, err := s.Pop()
            assert.NoError(t, err)
        }
    })
}