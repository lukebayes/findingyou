package display

import (
	"assert"
	"clock"
	"github.com/fogleman/ease"
	"testing"
	"time"
)

func TestTransition(t *testing.T) {

	var createTree = func() (Displayable, clock.FakeClock) {
		fakeClock := clock.NewFake()
		root, _ := Box(NewBuilderUsing(fakeClock), Children(func(b Builder) {
			moveRight := Transition(b,
				X,
				100.0,
				200.0,
				200,
				ease.Linear)
			Box(b, ID("abcd"), moveRight, ExcludeFromLayout(true))
		}))

		return root, fakeClock
	}

	t.Run("Instantiable", func(t *testing.T) {
		root, fakeClock := createTree()
		// Begin listening for enter frame events
		defer root.Builder().Destroy()
		go root.Builder().Listen()

		child := root.ChildAt(0)
		root.Layout()

		assert.Equal(t, int(child.X()), 100)
		// I expect enter frames to fire when this happens!
		// But they don't because they're currently implemented by the NanoWindow
		fakeClock.Add(101 * time.Millisecond)
		assert.Equal(t, int(child.X()), 150)
		fakeClock.Add(51 * time.Millisecond)
		assert.Equal(t, int(child.X()), 175)
		fakeClock.Add(51 * time.Millisecond)
		assert.Equal(t, int(child.X()), 200)
		fakeClock.Add(51 * time.Millisecond)
		assert.Equal(t, int(child.X()), 200)
	})

	t.Run("Updateable", func(t *testing.T) {
		t.Skip()

		root, fakeClock := createTree()
		child := root.ChildAt(0)
		root.Layout()

		// Begin listening for enter frame events
		defer root.Builder().Destroy()
		go root.Builder().Listen()

		fakeClock.Add(51 * time.Millisecond)
		assert.Equal(t, int(child.X()), 125)
		root.InvalidateChildren()
		fakeClock.Add(51 * time.Millisecond)

		child = root.FindComponentByID("abcd")
		assert.Equal(t, int(child.X()), 150)
	})
}
