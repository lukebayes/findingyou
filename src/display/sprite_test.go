package display

import (
	"assert"
	"strings"
	"testing"
)

func TestSprite(t *testing.T) {

	t.Run("AddChild", func(t *testing.T) {
		root := NewSprite()
		one := NewSprite()
		two := NewSprite()
		root.Width(200)
		assert.Equal(root.AddChild(one), 1)
		assert.Equal(root.AddChild(two), 2)
		assert.Equal(one.GetParent().GetId(), root.GetId())
		assert.Equal(two.GetParent().GetId(), root.GetId())
		assert.Nil(root.GetParent())
	})

	t.Run("GetChildCount", func(t *testing.T) {
		root := NewSprite()
		one := NewSprite()
		two := NewSprite()
		three := NewSprite()
		root.AddChild(one)
		one.AddChild(two)
		one.AddChild(three)

		assert.Equal(root.GetChildCount(), 1)
		assert.Equal(root.GetChildAt(0), one)

		assert.Equal(one.GetChildCount(), 2)
		assert.Equal(one.GetChildAt(0), two)
		assert.Equal(one.GetChildAt(1), three)
	})

	t.Run("GetFilteredChildren", func(t *testing.T) {
		createTree := func() (Displayable, []Displayable) {

			root := NewSprite()
			one := NewSpriteWithOpts(&Opts{Id: "a-t-one"})
			two := NewSpriteWithOpts(&Opts{Id: "a-t-two"})
			three := NewSpriteWithOpts(&Opts{Id: "b-t-three"})
			four := NewSpriteWithOpts(&Opts{Id: "b-t-four"})

			root.AddChild(one)
			root.AddChild(two)
			root.AddChild(three)
			root.AddChild(four)

			return root, []Displayable{one, two, three, four}
		}

		allKids := func(d Displayable) bool {
			return strings.Index(d.GetId(), "-t-") > -1
		}

		bKids := func(d Displayable) bool {
			return strings.Index(d.GetId(), "b-") > -1
		}

		t.Run("returns Empty slice", func(t *testing.T) {
			root := NewSprite()
			filtered := root.GetFilteredChildren(allKids)
			assert.Equal(len(filtered), 0)
		})

		t.Run("returns all matched children in simple match", func(t *testing.T) {
			root, _ := createTree()
			filtered := root.GetFilteredChildren(allKids)
			assert.Equal(len(filtered), 4)
		})

		t.Run("returns all matched children in harder match", func(t *testing.T) {
			root, _ := createTree()
			filtered := root.GetFilteredChildren(bKids)
			assert.Equal(len(filtered), 2)
			assert.Equal(filtered[0].GetId(), "b-t-three")
			assert.Equal(filtered[1].GetId(), "b-t-four")
		})
	})

	t.Run("GetChildren returns empty list", func(t *testing.T) {
		root := NewSprite()
		children := root.GetChildren()

		if children == nil {
			t.Error("GetChildren should not return nil")
		}

		assert.Equal(len(children), 0)
	})

	t.Run("GetChildren returns new list", func(t *testing.T) {
		root := NewSprite()
		one := NewSprite()
		two := NewSprite()
		three := NewSprite()

		root.AddChild(one)
		root.AddChild(two)
		root.AddChild(three)

		children := root.GetChildren()
		assert.Equal(len(children), 3)
	})
}
