package controls

import (
	"assert"
	"ui/context"
	"testing"
)

func TestFps(t *testing.T) {
	t.Run("Instantiable", func(t *testing.T) {
		instance := FPS(context.New())
		assert.NotNil(t, instance)
	})
}