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

func From11() {
	fmt.Println("I'm from V1.1.0")
}
func From12() {
	fmt.Println("I'm from V1.2.0")
}

func From13() {
	fmt.Println("I'm from V1.3.0")
}
