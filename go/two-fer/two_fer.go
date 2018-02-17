package twofer

import "fmt"

// ShareWith should have a comment documenting it.
func ShareWith(you string) string {
	if you == "" {
		you = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", you)
}
