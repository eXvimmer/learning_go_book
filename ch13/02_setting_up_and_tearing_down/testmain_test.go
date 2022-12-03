package testmain

import (
	"fmt"
	"os"
	"testing"
	"time"
)

// WARN: don't use package level variables in your program
var testTime time.Time

// TestMain will be invoked once to run stuff before all and after all tests
func TestMain(m *testing.M) {
	// configure state before all
	fmt.Println("Setting up general state before running tests")
	testTime = time.Now()
	// run all tests
	exitVal := m.Run() // 0 means all tests passed
	// after all
	fmt.Println("Clean up stuff after tests here")
	os.Exit(exitVal)
}

// NOTE: you can have only one TestMain per package

func TestFirst(t *testing.T) {
	fmt.Println("TestFirst uses stuff set up in TestMain", testTime)
}

func TestSecond(t *testing.T) {
	fmt.Println("TestSecond uses stuff set up in TestMain", testTime)
}
