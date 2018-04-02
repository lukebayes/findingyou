package display

type ComponentOption (func(d Displayable) error)

// ID will set the Component.Id.
func ID(value string) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().ID = value
		return nil
	}
}

// Title will set Component.Title.
func Title(value string) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().Title = value
		return nil
	}
}

// ExcludeFromLayout will configure Component.ExcludeFromLayout.
func ExcludeFromLayout(value bool) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().ExcludeFromLayout = value
		return nil
	}
}

// ActualWidth will set Component.ActualWidth.
func ActualWidth(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().ActualWidth = value
		return nil
	}
}

// ActualHeight will set Component.ActualHeight.
func ActualHeight(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().ActualHeight = value
		return nil
	}
}

// Width will set Component.Width.
func Width(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().Width = value
		return nil
	}
}

// Height will set Component.Height.
func Height(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().Height = value
		return nil
	}
}

// Size will set Component.Width and Component.Height.
func Size(width, height float64) ComponentOption {
	return func(d Displayable) error {
		model := d.GetModel()
		model.Width = width
		model.Height = height
		return nil
	}
}

// MaxWidth will set Component.MaxWidth.
func MaxWidth(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().MaxWidth = value
		return nil
	}
}

// MaxHeight will set Component.MaxHeight.
func MaxHeight(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().MaxHeight = value
		return nil
	}
}

// MinWidth will set Component.MinWidth.
func MinWidth(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().MinWidth = value
		return nil
	}
}

// MinHeight will set Component.MinHeight.
func MinHeight(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().MinHeight = value
		return nil
	}
}

// PrefWidth will set Component.PrefWidth.
func PrefWidth(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().PrefWidth = value
		return nil
	}
}

// PrefHeight will set Component.PrefHeight.
func PrefHeight(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().PrefHeight = value
		return nil
	}
}

// FlexWidth will set Component.FlexWidth.
func FlexWidth(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().FlexWidth = value
		return nil
	}
}

// FlexHeight will set Component.FlexHeight.
func FlexHeight(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().FlexHeight = value
		return nil
	}
}

// HAlign will set Component.HAlign.
func HAlign(align Alignment) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().HAlign = align
		return nil
	}
}

// VAlign will set Component.VAlign.
func VAlign(align Alignment) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().VAlign = align
		return nil
	}
}

// X will set Component.X.
func X(pos float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().X = pos
		return nil
	}
}

// Y will set Component.Y.
func Y(pos float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().Y = pos
		return nil
	}
}

// Z will set Component.Z.
func Z(pos float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().Z = pos
		return nil
	}
}

// LayoutType will set Component.LayoutType.
func LayoutType(layoutType LayoutTypeValue) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().LayoutType = layoutType
		return nil
	}
}

// Padding will set Component.Padding, which will effectively set padding for
// all four sides as well (bottom, top, left, right, horizontal and vertical).
func Padding(value float64) ComponentOption {
	return func(d Displayable) error {
		d.Padding(value)
		return nil
	}
}

// PaddingBottom will set Component.PaddingBottom.
func PaddingBottom(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().PaddingBottom = value
		return nil
	}
}

// PaddingLeft will set Component.PaddingLeft.
func PaddingLeft(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().PaddingLeft = value
		return nil
	}
}

// PaddingRight will set Component.PaddingRight.
func PaddingRight(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().PaddingRight = value
		return nil
	}
}

// PaddingTop will set Component.PaddingTop.
func PaddingTop(value float64) ComponentOption {
	return func(d Displayable) error {
		d.GetModel().PaddingTop = value
		return nil
	}
}

// AttrStyles will set Component.AttrStyles.
func AttrStyles(opts ...StyleOption) ComponentOption {
	styles := NewStyleDefinition()
	return func(d Displayable) error {
		for _, opt := range opts {
			opt(styles)
		}
		d.Styles(styles)
		return nil
	}
}

// Children will compose child components onto the current component by
// providing a closure that either accepts zero arguments, or accepts a single
// argument which will be a function that, when called will invalidate the
// component instance for a future render.
func Children(composer interface{}) ComponentOption {
	return func(d Displayable) error {
		return d.Composer(composer)
	}
}
