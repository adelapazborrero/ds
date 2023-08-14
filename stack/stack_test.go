package stack

import (
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	tests := []struct {
		name string
		want *Stack[int]
	}{
		{
			name: "Create a new integer stack",
			want: &Stack[int]{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack[int](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Add(t *testing.T) {
	tests := []struct {
		name      string
		setup     func(s *Stack[int])
		wantItems []int
	}{
		{
			name: "Add one element",
			setup: func(s *Stack[int]) {
				s.Add(1)
			},
			wantItems: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStack[int]()
			tt.setup(s)

			if !reflect.DeepEqual(s.Items, tt.wantItems) {
				t.Errorf("Stack items = %v, want %v", s.Items, tt.wantItems)
			}
		})
	}
}

func TestStack_Remove(t *testing.T) {
	tests := []struct {
		name      string
		setup     func(s *Stack[int])
		wantItems []int
	}{
		{
			name: "Remove one element",
			setup: func(s *Stack[int]) {
				s.Add(1)
				s.Remove()
			},
			wantItems: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStack[int]()
			tt.setup(s)

			if !reflect.DeepEqual(s.Items, tt.wantItems) {
				t.Errorf("Stack items = %v, want %v", s.Items, tt.wantItems)
			}
		})
	}
}
