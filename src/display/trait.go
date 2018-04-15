package display

import (
	"strings"
)

// Trait is a concrete factory function that builds a bag of ComponentOptions
// and applies them to all Selected Components before applying
// instance-specified options.
func Trait(b Builder, selector string, opts ...ComponentOption) error {
	component := b.Peek()
	if component == nil {
		panic("Trait definition must be nested inside of a component")
	}

	// TODO(lbayes): There are more questions here than answers.
	// This entire feature is not fleshed out at all and will certainly behave
	// unexpectedly.
	component.PushTrait(selector, opts...)
	return nil
}

func mergeSelectOptions(result, next TraitOptions) TraitOptions {
	for key, value := range next {
		result[key] = value
	}
	return result
}

func selectorMatches(key string, d Displayable) bool {
	// Return for the "all" selector
	if key == "*" {
		return true
	}

	// Return for TypeName match
	if key == d.TypeName() {
		return true
	}

	// Return for ID match
	if strings.Index(key, "#") == 1 {
		if d.ID() == key[1:len(key)-1] {
			return true
		}

		// Return for ID style, key but no match
		return false
	}

	// Return for any TraitName match
	for _, name := range d.TraitNames() {
		if key == name {
			return true
		}
	}

	// The provided selector does not match this instance.
	return false
}

func TraitOptionsFor(d, parent Displayable) []ComponentOption {
	optionsMap := d.TraitOptions()
	current := parent
	for current != nil {
		optionsMap = mergeSelectOptions(optionsMap, current.TraitOptions())
		current = current.Parent()
	}

	result := []ComponentOption{}
	for key, value := range optionsMap {
		if selectorMatches(key, d) {
			result = append(result, value...)
		}
	}

	return result
}
