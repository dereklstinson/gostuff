package testingtest

import "testing"

//To test Add run command "go test"

//TestAdd takes the type pointer testing.T and you will call the function add that I put in another file
func TestAdd(t *testing.T) { //Test_add worked for me too
	a, b := 1, 2
	if a+b != Add(a, b) {
		t.Error("no")

	}

}
