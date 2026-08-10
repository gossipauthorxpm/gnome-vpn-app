package common

import "github.com/diamondburned/gotk4-adwaita/pkg/adw"

func HasViewStackContainsName(stack *adw.ViewStack, name string) bool {
	return stack.ChildByName(name) != nil
}
