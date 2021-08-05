package variableswithinitializers

type MyStruct struct {
	i      int
	j      int
	c      bool
	python bool
	java   string
}

var i, j int = 1, 2

func VariablesWithInitializers() MyStruct {
	var c, python, java = true, false, "no!"
	// fmt.Println(i, j, c, python, java)
	// message := multipleresults.MultipleResults()
	// fmt.Println(message)

	v := MyStruct{i, j, c, python, java}

	return v
}
