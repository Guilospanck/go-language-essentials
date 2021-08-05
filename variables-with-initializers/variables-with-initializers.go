package variableswithinitializers

import (
	"fmt"

	multipleresults "github.com/Guilospanck/GoLanguageEssentials/multiple-results"
)

var i, j int = 1, 2

func VariablesWithInitializers() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	message := multipleresults.MultipleResults()
	fmt.Println(message)
}
