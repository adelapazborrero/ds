package queue

import (
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	tests := []struct {
		name string
		want *Queue[int]
	}{
		{
			name: "Create a new integer queue",
			want: &Queue[int]{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQueue[int](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Add(t *testing.T) {
	tests := []struct {
		name      string
		setup     func(q *Queue[int])
		wantItems []int
	}{
		{
			name: "Add one element",
			setup: func(q *Queue[int]) {
				q.Add(1)
			},
			wantItems: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue[int]()
			tt.setup(q)

			if !reflect.DeepEqual(q.Items, tt.wantItems) {
				t.Errorf("Queue items = %v, want %v", q.Items, tt.wantItems)
			}
		})
	}
}

func TestQueue_Remove(t *testing.T) {
	tests := []struct {
		name      string
		setup     func(q *Queue[int])
		wantItems []int
	}{
		{
			name: "Remove one element",
			setup: func(q *Queue[int]) {
				q.Add(1)
				q.Remove()
			},
			wantItems: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue[int]()
			tt.setup(q)

			if !reflect.DeepEqual(q.Items, tt.wantItems) {
				t.Errorf("Queue items = %v, want %v", q.Items, tt.wantItems)
			}
		})
	}
}
