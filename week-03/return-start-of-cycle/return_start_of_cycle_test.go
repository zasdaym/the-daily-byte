package returnstartofcycle

import (
	"testing"

	"github.com/zasdaym/the-daily-byte/linkedlist"
)

func TestStartOfCycle(t *testing.T) {
	t.Run("no cycle", func(t *testing.T) {
		head := linkedlist.New([]int{1, 2, 3, 4, 5})
		var want *linkedlist.Node
		if got := StartOfCycle(head); got != want {
			t.Errorf("ContainsCycle(%s) = %v, want %v", head, got, want)
		}
	})

	t.Run("with cycle", func(t *testing.T) {
		var (
			fourth = &linkedlist.Node{Value: 4}
			third  = &linkedlist.Node{Value: 3, Next: fourth}
			second = &linkedlist.Node{Value: 2, Next: third}
			first  = &linkedlist.Node{Value: 1, Next: second}
		)
		fourth.Next = first
		want := fourth
		got := StartOfCycle(first)
		if got == nil {
			t.Fatalf("got %v, want %d", got, want.Value)
		}
		if got.Value != want.Value {
			t.Fatalf("got %d, want %d", got.Value, want.Value)
		}
	})
}
