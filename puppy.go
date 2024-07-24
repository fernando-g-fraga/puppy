// git
package puppy

import (
	"github.com/fernando-g-fraga/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Barks())
}
func BigBarks() string {
	return dog.WhenGrownUp(Bark())
}
