// git
package puppy

import (
	"fmt"

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

func From1() {
	fmt.Println("I'm from V1.0.0")
}
