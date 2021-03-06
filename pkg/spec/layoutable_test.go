package spec_test

import (
	"testing"

	"github.com/waybeams/assert"
	"github.com/waybeams/waybeams/pkg/fakes"
	"github.com/waybeams/waybeams/pkg/opts"
	"github.com/waybeams/waybeams/pkg/spec"
)

func TestLayoutable(t *testing.T) {

	t.Run("ContentWidth", func(t *testing.T) {
		ctrl := fakes.Fake()
		ctrl.SetContentWidth(123)
		assert.Equal(ctrl.ContentWidth(), 123)
	})

	t.Run("ContentHeight", func(t *testing.T) {
		ctrl := fakes.Fake()
		ctrl.SetContentHeight(124)
		assert.Equal(ctrl.ContentHeight(), 124)
	})

	t.Run("Default Size", func(t *testing.T) {
		ctrl := fakes.Fake()
		assert.Equal(ctrl.Width(), 0, "Width")
		assert.Equal(ctrl.Height(), 0, "Height")
	})

	t.Run("LayoutType() default value", func(t *testing.T) {
		ctrl := fakes.Fake()
		assert.Equal(ctrl.LayoutType(), spec.NoLayoutType)
	})

	t.Run("MaxHeight constrained Height", func(t *testing.T) {
		ctrl := fakes.Fake(opts.Height(51), opts.MaxHeight(41))
		assert.Equal(ctrl.Height(), 41.0)
	})

	t.Run("MaxWidth constrained Width", func(t *testing.T) {
		ctrl := fakes.Fake(opts.Width(50), opts.MaxWidth(40))
		assert.Equal(ctrl.Width(), 40.0)
	})

	t.Run("MinHeight becomes unset Height", func(t *testing.T) {
		ctrl := fakes.Fake(opts.MinHeight(20))
		assert.Equal(ctrl.Height(), 20.0)
	})

	t.Run("MinWidth becomes unset Width", func(t *testing.T) {
		ctrl := fakes.Fake(opts.MinWidth(20))
		assert.Equal(ctrl.Width(), 20.0)
	})

	t.Run("MinHeight replaces existing Height", func(t *testing.T) {
		ctrl := fakes.Fake(opts.MinHeight(20), opts.Height(10))
		assert.Equal(ctrl.Height(), 20.0)
	})

	t.Run("MinWidth replaces existing Width", func(t *testing.T) {
		ctrl := fakes.Fake(opts.MinWidth(20), opts.Width(10))
		assert.Equal(ctrl.Width(), 20.0)
	})

	t.Run("PrefWidth default value", func(t *testing.T) {
		assert.Equal(fakes.Fake().PrefWidth(), 0)
	})

	t.Run("PrefWidth ControlModel value", func(t *testing.T) {
		ctrl := fakes.Fake(opts.PrefWidth(200))
		assert.Equal(ctrl.PrefWidth(), 200)
	})

	/*
		// These should only work after applying Stack or Flow Layouts!
		// Specs should not be this smart.

		t.Run("Child updates Min size", func(t *testing.T) {
			ctrl := fakes.FakeSpec(
				spec.Child(fakes.FakeSpec(opts.Width(35), opts.Height(55))),
				spec.Child(fakes.FakeSpec(opts.Width(50), opts.Height(30))),
			)

			assert.Equal(ctrl.Width(), 50)
			assert.Equal(ctrl.Height(), 55)
		})

		t.Run("WidthInBounds from Child expansion plus Padding", func(t *testing.T) {
			ctrl := fakes.FakeSpec(spec.Padding(10), opts.Width(30), opts.Height(20),
				spec.Child(fakes.FakeSpec(opts.MinWidth(50), opts.MinHeight(40))),
				spec.Child(fakes.FakeSpec(opts.MinWidth(30), opts.MinHeight(30))),
			)

			assert.Equal(ctrl.Width(), 50.0)
			assert.Equal(ctrl.Height(), 40.0)
		})
	*/

	/*

		t.Run("GetOffsetFor", func(t *testing.T) {
			t.Run("Root at 0,0", func(t *testing.T) {
				root := Box(context.New())
				xOffset := root.XOffset()
				yOffset := root.YOffset()
				assert.Equal(xOffset, 0)
				assert.Equal(yOffset, 0)
			})

			t.Run("Root at offset", func(t *testing.T) {
				root := Box(context.New(), X(10), Y(15))
				xOffset := root.XOffset()
				yOffset := root.YOffset()
				assert.Equal(xOffset, 10)
				assert.Equal(yOffset, 15)
			})

			t.Run("Child receives offset for padding", func(t *testing.T) {
				var root, child Displayable
				root = Box(context.New(), Padding(10), Width(100), Height(100), Children(func(c Context) {
					child = Box(c, FlexWidth(1), FlexHeight(1))
				}))

				assert.Equal(root.XOffset(), 0)
				assert.Equal(child.XOffset(), 10)
			})

			t.Run("Child at double offset", func(t *testing.T) {
				var nestedChild Displayable
				// TODO(lbayes): Possible inadvertent duplication during large refactoring, but if test is failing, fix indentation instead?
				// NOSUBMIT DELETE THIS COMMENT WHEN TESTS PASSING
				// Box(context.New(), Padding(10), Children(func(c Context) {
				Box(context.New(), Padding(10), Children(func(c Context) {
					Box(c, Padding(15), Children(func() {
						nestedChild = Box(c, Padding(10))
					}))
				}))

				xOffset := nestedChild.XOffset()
				yOffset := nestedChild.YOffset()
				assert.Equal(xOffset, 25)
				assert.Equal(yOffset, 25)
			})
		})

		t.Run("Padding", func(t *testing.T) {
			t.Run("DefaultPadding", func(t *testing.T) {
				box := Box(context.New())

				assert.Equal(box.Padding(), -1, "Default Padding")
				assert.Equal(box.PaddingBottom(), -1, "Default PaddingBottom")
				assert.Equal(box.PaddingTop(), -1, "Default PaddingTop")
				assert.Equal(box.PaddingLeft(), -1, "Default PaddingLeft")
				assert.Equal(box.PaddingRight(), -1, "Default PaddingRight")

				assert.Equal(box.MinWidth(), -1, "GetMinWidth")
				assert.Equal(box.MinHeight(), -1, "GetMinWidth")

				assert.Equal(box.Width(), 0, "Width")
			})

			t.Run("Override side padding", func(t *testing.T) {
				box := Box(context.New(), Padding(10))
				assert.Equal(box.Padding(), 10, "Default Padding")
				assert.Equal(box.PaddingBottom(), 10, "Default PaddingBottom")
				assert.Equal(box.PaddingTop(), 10, "Default PaddingTop")
				assert.Equal(box.PaddingLeft(), 10, "Default PaddingLeft")
				assert.Equal(box.PaddingRight(), 10, "Default PaddingRight")
			})

			t.Run("Interacts with GetMinWidth()", func(t *testing.T) {
				box := Box(context.New(), Padding(10))
				assert.Equal(box.MinWidth(), 20, "GetMinWidth")
				assert.Equal(box.MinHeight(), 20, "GetMinWidth")
			})

			t.Run("Applying Padding spreads to all four sides", func(t *testing.T) {
				root := TestControl(context.New(), Padding(10))

				assert.Equal(root.HorizontalPadding(), 20.0)
				assert.Equal(root.VerticalPadding(), 20.0)

				assert.Equal(root.PaddingBottom(), 10.0)
				assert.Equal(root.PaddingLeft(), 10.0)
				assert.Equal(root.PaddingRight(), 10.0)
				assert.Equal(root.PaddingTop(), 10.0)
			})

			t.Run("PaddingTop overrides Padding", func(t *testing.T) {
				root := TestControl(context.New(), Padding(10), PaddingTop(5))
				assert.Equal(root.PaddingTop(), 5.0)
				assert.Equal(root.PaddingBottom(), 10.0)
				assert.Equal(root.Padding(), 10.0)
			})

			t.Run("PaddingTop overrides Padding regardless of order", func(t *testing.T) {
				root := TestControl(context.New(), PaddingTop(5), Padding(10))
				assert.Equal(root.PaddingTop(), 5.0)
				assert.Equal(root.PaddingBottom(), 10.0)
				assert.Equal(root.Padding(), 10.0)
			})
		})
	*/
}
