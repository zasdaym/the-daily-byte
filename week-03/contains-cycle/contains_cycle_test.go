package containscycle

import (
	"testing"

	"github.com/zasdaym/the-daily-byte/linkedlist"
)

func TestContainsCycle(t *testing.T) {
	t.Run("no cycle", func(t *testing.T) {
		head := linkedlist.New([]int{1, 2, 3, 4, 5})
		want := false
		if got := ContainsCycle(head); got != want {
			t.Errorf("ContainsCycle(%s) = %v, want %v", head, got, want)
		}
	})

	t.Run("with cycle", func(t *testing.T) {
		var (
			third  = &linkedlist.Node{Value: 3}
			second = &linkedlist.Node{Value: 2, Next: third}
			first  = &linkedlist.Node{Value: 1, Next: second}
		)
		third.Next = first
		want := true
		if got := ContainsCycle(first); got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
