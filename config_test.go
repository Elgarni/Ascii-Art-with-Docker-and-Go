package main

import (
	"os"
	"testing"
)

/*
Since some of the tests mutates the underlying OS env vars,
it is recommended that the tests are run in a disposable container.
*/


func TestReadingEnvVarsProvidedCorrectTextAndMove(t *testing.T)  {
	// prepare OS env vars
	expectedText := "myText"
	expectedMove := "blink"
	os.Setenv("text", expectedText)
	os.Setenv("move", expectedMove)

	// read OS env vars, and test them
	envConfig := readEnvVars()
	assertEqual(t, envConfig.text, expectedText)
	assertEqual(t, envConfig.moveType, expectedMove)
}


func TestReadingEnvVarsProvidedCorrectTBase64EncodedextAndMove(t *testing.T)  {
	// prepare OS env vars
	expectedText := "Good morning"
	expectedMove := "blink"
	os.Setenv("text", "R29vZCBtb3JuaW5n")
	os.Setenv("move", expectedMove)
	os.Setenv("hidden", "true")

	// read OS env vars, and test them
	envConfig := readEnvVars()
	assertEqual(t, envConfig.text, expectedText)
	assertEqual(t, envConfig.moveType, expectedMove)
}


// test reading the env vars when none is set
func TestReadingEnvVarsProvidedNone(t *testing.T) {

	// unset any expected env vars
	os.Unsetenv("text")
	os.Unsetenv("move")
	os.Unsetenv("t")
	os.Unsetenv("m")

	// read OS env vars, and test them
	envConfig := readEnvVars()
	assertEqual(t, envConfig.moveType, "scroll")
	assertNumberGreaterThan(t, len(envConfig.text), 0);
	assertNumberGreaterThan(t, envConfig.duration, 0);
}



// asserts if two values are equal
func assertEqual(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Fatalf("Expected %s, but found %s", expected, actual)
	}
}

// asserts if a number is greater than a lower bound
func assertNumberGreaterThan(t *testing.T, number int, lowerLimit int)  {
	if number <= lowerLimit {
		t.Fatalf("Number %d is not greater than %d", number, lowerLimit)
	}
}